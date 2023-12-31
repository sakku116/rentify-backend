package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type EnvsSchema struct {
	HOST          string
	PORT          int
	JWT_SECRET    string
	JWT_EXP       int
	MONGO_URI     string
	MONGO_DB_NAME string
}

var Envs *EnvsSchema

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	Envs = &EnvsSchema{
		HOST:          viper.GetString("HOST"),
		PORT:          viper.GetInt("PORT"),
		JWT_SECRET:    viper.GetString("JWT_SECRET"),
		JWT_EXP:       viper.GetInt("JWT_EXP"),
		MONGO_URI:     viper.GetString("MONGO_URI"),
		MONGO_DB_NAME: viper.GetString("MONGO_DB_NAME"),
	}
}
