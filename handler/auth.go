package handler

import (
	"rentify/domain/rest"
	"rentify/service"
	"rentify/utils/http_response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	respWriter  http_response.IResponseWriter
	authService service.AuthService
}

func NewAuthHandler(respWriter http_response.IResponseWriter, authService service.AuthService) AuthHandler {
	return AuthHandler{
		respWriter:  respWriter,
		authService: authService,
	}
}

func (slf *AuthHandler) Login(ctx *gin.Context) {
	var payload rest.PostLoginReq
	err := ctx.BindJSON(&payload)
	if err != nil {
		slf.respWriter.HTTPJsonErr(ctx, 400, "bad request", err.Error(), nil)
		return
	}

	token, err := slf.authService.Login(ctx, payload.Username, payload.Password)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	slf.respWriter.HTTPJson(ctx, gin.H{
		"access_token": token,
	})
}

func (slf *AuthHandler) CheckToken(ctx *gin.Context) {
	var payload rest.PostCheckTokenReq
	err := ctx.BindJSON(&payload)
	if err != nil {
		slf.respWriter.HTTPJsonErr(ctx, 400, "bad request", err.Error(), nil)
		return
	}

	user, err := slf.authService.CheckToken(ctx, payload.Token)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	slf.respWriter.HTTPJson(ctx, gin.H{
		"username": user.Username,
		"role":     user.Role,
	})
}

func (slf *AuthHandler) SetRoleFromToken(ctx *gin.Context) {
	var payload rest.PostSetRoleFromTokenReq
	err := ctx.BindJSON(&payload)
	if err != nil {
		slf.respWriter.HTTPJsonErr(ctx, 400, "bad request", err.Error(), nil)
		return
	}

	// check token
	user, err := slf.authService.CheckToken(ctx, payload.Token)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	// set role
	err = slf.authService.SetRole(ctx, user.ID, payload.Role)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	// regenerate token
	token, err := slf.authService.Login(ctx, user.Username, user.Password)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	slf.respWriter.HTTPJson(ctx, gin.H{
		"access_token": token,
		"role":         user.Role,
		"username":     user.Username,
	})
}

func (slf *AuthHandler) Register(ctx *gin.Context) {
	var payload rest.PostRegisterReq
	err := ctx.BindJSON(&payload)
	if err != nil {
		slf.respWriter.HTTPJsonErr(ctx, 400, "bad request", err.Error(), nil)
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
		slf.respWriter.HTTPJsonErr(ctx, 400, "all fields must be filled", "", nil)
		return
	}

	// match password validation
	if payload.Password != payload.ConfirmPassword {
		slf.respWriter.HTTPJsonErr(ctx, 400, "password and confirm password must be same", "", nil)
		return
	}

	err = slf.authService.Register(ctx, payload.Username, payload.Email, payload.Password)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	// regenerate token
	token, err := slf.authService.Login(ctx, payload.Username, payload.ConfirmPassword)
	if err != nil {
		slf.respWriter.HTTPCustomErr(ctx, err)
		return
	}

	slf.respWriter.HTTPJson(ctx, gin.H{
		"access_token": token,
	})
}
