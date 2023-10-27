package handler

import (
	"fmt"
	"rentify/dto"
	"rentify/exception"
	"rentify/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

func (slf *AuthHandler) Login(ctx *gin.Context) {
	var payload dto.PostLoginReq
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	token, err := slf.authService.Login(ctx, payload.Username, payload.Password)
	if err != nil {
		switch err {
		case exception.AuthUserPassRequired:
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "username and password are required",
			})
		case exception.DbObjNotFound:
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": fmt.Sprintf("username %s not found", payload.Username),
			})
		case exception.AuthPasswordIncorrect:
			ctx.JSON(401, gin.H{
				"error":   true,
				"message": err.Error(),
			})
		default:
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": err.Error(),
			})
		}
		return
	}

	ctx.JSON(200, gin.H{
		"error":        false,
		"message":      "OK",
		"access_token": token,
	})
}

func (slf *AuthHandler) CheckToken(ctx *gin.Context) {
	var payload dto.PostCheckTokenReq
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	user, err := slf.authService.CheckToken(ctx, payload.Token)
	if err != nil {
		switch err {
		case exception.AuthInvalidToken, exception.AuthUserNotFound:
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "invalid token",
			})
		case exception.AuthUserBanned:
			ctx.JSON(403, gin.H{
				"error":   true,
				"message": "user is banned",
			})
		default:
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": err.Error(),
			})
		}
		return
	}

	// check for role
	if user.Role == "" {
		ctx.JSON(403, gin.H{
			"error":   true,
			"message": "role is not set",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "OK",
		"data": gin.H{
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

func (slf *AuthHandler) SetRoleFromToken(ctx *gin.Context) {
	var payload dto.PostSetRoleFromTokenReq
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	// check token
	user, err := slf.authService.CheckToken(ctx, payload.Token)
	if err != nil {
		switch err {
		case exception.AuthInvalidToken, exception.AuthUserNotFound:
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "invalid token",
			})
		case exception.AuthUserBanned:
			ctx.JSON(403, gin.H{
				"error":   true,
				"message": "user is banned",
			})
		default:
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": err.Error(),
			})
		}
		return
	}

	// set role
	err = slf.authService.SetRole(ctx, user.ID, payload.Role)
	if err != nil {
		switch err {
		case exception.AuthInvalidRole:
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "invalid role",
			})
			return
		default:
			ctx.JSON(500, gin.H{
				"error":   true,
				"message": err.Error(),
			})
		}
	}

	// regenerate token
	token, err := slf.authService.Login(ctx, user.Username, user.Password)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"error":        false,
		"message":      "OK",
		"access_token": token,
		"data": gin.H{
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

func (slf *AuthHandler) Register(ctx *gin.Context) {
	var payload dto.PostRegisterReq
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	// validate all field must be filled
	if payload.Username == "" ||
		payload.Password == "" ||
		payload.Email == "" ||
		payload.ConfirmPassword == "" {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "username and password are required",
		})
		return
	}

	// match password validation
	if payload.Password != payload.ConfirmPassword {
		ctx.JSON(400, gin.H{
			"error":   true,
			"message": "password and confirm password must be same",
		})
		return
	}

	err = slf.authService.Register(ctx, payload.Username, payload.Email, payload.Password)
	if err != nil {
		switch err {
		case exception.UserAlreadyExistByEmail:
			ctx.JSON(
				400, gin.H{
					"error":   true,
					"message": fmt.Sprintf("user with email %s already exist", payload.Email),
				},
			)
		case exception.UserAlreadyExistByUsername:
			ctx.JSON(
				400, gin.H{
					"error":   true,
					"message": fmt.Sprintf("user with username %s already exist", payload.Username),
				},
			)
		default:
			ctx.JSON(
				500, gin.H{
					"error":   true,
					"message": err.Error(),
				},
			)
		}
		return
	}

	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "OK",
	})
}
