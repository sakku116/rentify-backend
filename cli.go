package main

import (
	"context"
	"fmt"
	"os"
	"rentify/cli"
	"rentify/config"
	"rentify/repository"
)

func CliHandler(args []string) {
	args = args[1:]
	ctx := context.Background()

	mongoConn := config.NewMongoConn(ctx)
	defer mongoConn.Close(ctx)
	mongoDB := mongoConn.Database("rentify")

	userRepo := repository.NewUserRepo(mongoDB.Collection("users"))

	switch args[0] {
	case "seed-superuser":
		fmt.Println("running seed superuser...")
		cli.SeedSuperuser(userRepo, args[1:]...)
	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}
	fmt.Println("done")
}
