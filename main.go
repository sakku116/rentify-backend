package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("starting rest api app...")
	log.Printf("Envs: %v", Envs)

	router := gin.Default()
	SetupRouter(router)
	router.Run(Envs.ADDR)
	fmt.Println("starting rest api app...")
}
