package handler

import (
	"rentify/domain/rest"
	"rentify/service"
	"rentify/utils/http_response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	respWriter  http_response.IResponseWriter
	authService service.IAuthService
}

type IAuthHandler interface {
	Login(ctx *gin.Context)
	CheckToken(ctx *gin.Context)
	SetRoleFromToken(ctx *gin.Context)
	Register(ctx *gin.Context)
}

func NewAuthHandler(respWriter http_response.IResponseWriter, authService service.IAuthService) AuthHandler {
	return AuthHandler{
		respWriter:  respWriter,
		authService: authService,
	}
}

// Login
// @Summary generate jwt token
// @Tags Auth
// @Success 200 {object} rest.BaseJSONResp{data=rest.AuthLoginResp}
// @Router /auth/login [post]
// @param payload  body  rest.PostLoginReq  true "payload"
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

	slf.respWriter.HTTPJson(ctx, rest.AuthLoginResp{
		AccessToken: token,
	})
}

// CheckToken
// @Summary check jwt token
// @Tags Auth
// @Success 200 {object} rest.BaseJSONResp{data=rest.AuthLoginResp}
// @Router /auth/check-token [post]
// @Security JWTAuth
// @param payload  body  rest.PostCheckTokenReq  true "payload"
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

	slf.respWriter.HTTPJson(ctx, rest.AuthCheckTokenResp{
		Username: user.Username,
		Role:     user.Role,
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
