package middleware

import (
	"rentify/service"
	"rentify/utils/http_response"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(respWriter http_response.IResponseWriter, authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			respWriter.HTTPJsonErr(c, 401, "Authorization header is missing", "", nil)
			c.Abort()
			return
		}

		token := strings.Split(tokenString, " ")[1]

		user, err := authService.CheckToken(c, token)
		if err != nil {
			respWriter.HTTPCustomErr(c, err)
			c.Abort()
			return
		}

		// check for role
		if user.Role == "" {
			respWriter.HTTPJsonErr(c, 403, "role is not set", "", nil)
			c.Abort()
			return
		}

		// Pass the claims to subsequent handlers
		c.Set("user_id", user.ID)
		c.Set("role", user.Role)
		c.Set("username", user.Username)

		c.Next()
	}
}
