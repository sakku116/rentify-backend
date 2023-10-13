package main


import (
	"github.com/spf13/viper"
)

func GetEnv() EnvsSchema {
	return EnvsSchema{
		ADDR: viper.GetString("ADDR"),
	}
}
