package main

import (
	"rentify/schemas"

	"github.com/spf13/viper"
)

func GetEnv() schemas.Envs {
	return schemas.Envs{
		ADDR: viper.GetString("ADDR"),
	}
}
