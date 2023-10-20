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

	mongoConn := config.NewMongoConn(context.Background())
	mongoDB := mongoConn.Database("rentify")
	defer mongoConn.Close(ctx)

	userRepo := repository.NewUserRepo(mongoDB.Collection("users"))

	switch args[0] {
	case "seed-superuser":
		fmt.Println("running seed superuser...")
		cli.SeedSuperuser(userRepo)
	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}
	fmt.Println("done")
}
