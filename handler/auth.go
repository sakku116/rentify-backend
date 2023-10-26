package handler

import (
	"fmt"
	"rentify/dto"
	"rentify/exception"
	"rentify/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (slf *AuthHandler) Login() (ctx *gin.Context) {
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
	return
}

func (slf *AuthHandler) CheckToken() (ctx *gin.Context) {
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

	ctx.JSON(200, gin.H{
		"error":    false,
		"message":  "OK",
		"username": user.Username,
	})
	return
}
