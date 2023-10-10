package main

import (
	"go-rest-api/schemas"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Envs schemas.Envs

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	log.Println("loading .env file to Envs variable")
	Envs = GetEnv()
}
