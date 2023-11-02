package main

import (
	"fmt"
	"log"
	"os"
	"rentify/config"

	"github.com/gin-gonic/gin"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		CliHandler(args)
	} else {
		log.Printf("Envs: %v", config.Envs)
		log.Println("starting rest api app...")

		router := gin.Default()
		SetupServer(router)
		router.Run(config.Envs.ADDR)

		fmt.Println("starting rest api app...")
	}

}
