package main

import (
	"context"
	"rentify/config"
	"rentify/handler"
	"rentify/middleware"
	"rentify/repository"
	"rentify/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	ctx := context.Background()
	mongoConn := config.NewMongoConn(ctx)
	// defer mongoConn.Close(ctx)

	mongoDb := mongoConn.Database("rentify")

	// repositories
	userRepo := repository.NewUserRepo(mongoDb.Collection("users"))

	// services
	authService := service.NewAuthService(userRepo)

	// handlers
	authHandler := handler.NewAuthHandler(authService)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error":   false,
			"message": "pong!",
		})
	})

	router.POST("/auth/login", authHandler.Login)
	router.POST("/auth/check-token", authHandler.CheckToken)
	router.POST("/auth/set-role", authHandler.SetRoleFromToken)
	router.POST("/auth/register", authHandler.Register)

	secureRouter := router.Group("/")
	{
		secureRouter.Use(middleware.JWTMiddleware(authService))

	}
}
