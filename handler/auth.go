package handler

import (
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

	token, err := slf.authService.Login(payload.Username, payload.Password)
	if err != nil {
		switch err {
		case exception.DbObjNotFound:
			ctx.JSON(400, gin.H{
				"error":   true,
				"message": "user not found",
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
