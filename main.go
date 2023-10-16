package main

import (
	"fmt"
	"log"
	"rentify/config"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("starting rest api app...")
	log.Printf("Envs: %v", config.Envs.ADDR)

	router := gin.Default()
	SetupRouter(router)
	router.Run(config.Envs.ADDR)
	fmt.Println("starting rest api app...")
}
