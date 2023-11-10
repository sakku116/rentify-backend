package main

import (
	"context"
	"rentify/config"
	"rentify/handler"
	"rentify/middleware"
	"rentify/repository"
	"rentify/service"
	"rentify/utils/http_response"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupServer(router *gin.Engine) {
	ctx := context.Background()
	mongoConn := config.NewMongoConn(ctx)

	mongoDb := mongoConn.Database("rentify")
	responseWriter := http_response.NewResponseWriter()

	// repositories
	userRepo := repository.NewUserRepo(mongoDb.Collection("users"))

	// services
	authService := service.NewAuthService(userRepo)

	// handlers
	authHandler := handler.NewAuthHandler(responseWriter, authService)
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
		secureRouter.Use(middleware.JWTMiddleware(responseWriter, authService))
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
