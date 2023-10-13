package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type EnvsSchema struct {
	ADDR string
}

var Envs EnvsSchema

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	Envs = GetEnv()
}
