package main

import (
	"fmt"
	"log"
	"os"
	"rentify/config"
	_ "rentify/docs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @title Rentify API
// @securitydefinitions.apiKey BearerAuth
// @in header
// @name Authorization
// @default Bearer {token}
func main() {
	args := os.Args
	if len(args) > 1 {
		CliHandler(args)
	} else {
		log.Printf("Envs: %v", config.Envs)
		log.Println("starting rest api app...")

		router := gin.Default()
		SetupServer(router)
		router.Run(config.Envs.HOST + ":" + strconv.Itoa(config.Envs.PORT))

		fmt.Println("starting rest api app...")
	}
}
