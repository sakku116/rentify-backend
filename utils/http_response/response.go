package http_response

import (
	error_utils "rentify/utils/error"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct{}

type IResponseWriter interface {
	HTTPError(ctx *gin.Context, err error)
	HTTPJSON(ctx *gin.Context, code int, message string, detail string, data interface{})
}

func NewResponseWriter() IResponseWriter {
	return &ResponseWriter{}
}

func (r *ResponseWriter) HTTPError(ctx *gin.Context, err error) {
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

func (r *ResponseWriter) HTTPJSON(ctx *gin.Context, code int, message string, detail string, data interface{}) {
	ctx.JSON(code, gin.H{
		"error":   false,
		"message": message,
		"detail":  detail,
		"data":    data,
	})
}
