package middleware

import (
	"rentify/exception"
	"rentify/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(401, gin.H{
				"error":   true,
				"message": "Authorization header is missing",
			})
			c.Abort()
			return
		}

		token := strings.Split(tokenString, " ")[1]

		user, err := authService.CheckToken(c, token)
		if err != nil {
			switch err {
			case exception.AuthInvalidToken, exception.AuthUserNotFound:
				c.JSON(400, gin.H{
					"error":   true,
					"message": "invalid token",
				})
			case exception.AuthUserBanned:
				c.JSON(403, gin.H{
					"error":   true,
					"message": "user is banned",
				})
			default:
				c.JSON(500, gin.H{
					"error":   true,
					"message": err.Error(),
				})
			}
			c.Abort()
			return
		}

		// check for role
		if user.Role == "" {
			c.JSON(403, gin.H{
				"error":   true,
				"message": "role is not set",
			})
			return
		}

		// Pass the claims to subsequent handlers
		c.Set("user_id", user.ID)
		c.Set("role", user.Role)
		c.Set("username", user.Username)

		c.Next()
	}
}
