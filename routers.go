package main

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "pong!",
		})
	})

	router.POST("/auth/generate-token")
	router.GET("/auth/check-token")
}
