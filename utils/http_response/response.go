package http_response

import (
	error_utils "rentify/utils/error"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct{}

type IResponseWriter interface {
	HTTPCustomErr(ctx *gin.Context, err error)
	HTTPJsonErr(ctx *gin.Context, code int, message string, detail string, data interface{})
	HTTPJson(ctx *gin.Context, data interface{})
}

func NewResponseWriter() IResponseWriter {
	return &ResponseWriter{}
}

func (r *ResponseWriter) HTTPCustomErr(ctx *gin.Context, err error) {
	customErr, ok := err.(*error_utils.CustomErr)
	if ok {
		ctx.JSON(customErr.Code, gin.H{
			"error":   true,
			"message": customErr.Error(),
			"detail":  "",
			"data":    nil,
		})
		return
	}
	ctx.JSON(500, gin.H{
		"error":   true,
		"message": "internal server error",
		"detail":  err.Error(),
		"data":    nil,
	})
}

func (r *ResponseWriter) HTTPJsonErr(ctx *gin.Context, code int, message string, detail string, data interface{}) {
	ctx.JSON(code, gin.H{
		"error":   true,
		"message": message,
		"detail":  detail,
		"data":    data,
	})
}

func (r *ResponseWriter) HTTPJson(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"error":   false,
		"message": "OK",
		"detail":  "",
		"data":    data,
	})
}
