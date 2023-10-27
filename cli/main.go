package cli

import (
	"context"
	"fmt"
	"rentify/entity"
	"rentify/exception"
	"rentify/helper"
	"rentify/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// args should be empty for default seed (superuser;superuser@gmail.com;superuser)
// or args must be containing 3 strings for custom username, email and passwords
func SeedSuperuser(userRepo repository.UserRepo, args ...string) {
	ctx := context.Background()

	// validate args
	if len(args) != 3 && len(args) != 0 {
		fmt.Println("invalid args, should be empty for default seed (superuser;superuser@gmail.com;superuser) or 3 strings for custom username, email and passwords")
		return
	}

	username := "superuser"
	email := "superuser@gmail.com"
	password := "superuser"
	if len(args) == 3 {
		username = args[0]
		email = args[1]
		password = args[2]
	}

	// check for existing user
	existingUser, err := userRepo.GetByUsername(ctx, username)
	if err != nil && err != exception.DbObjNotFound {
		panic(err)
	}
	if existingUser != nil {
		fmt.Printf("%s already exists\n", username)
		return
	}

	uuid := helper.GenerateUUID()
	timeNow := time.Now().Unix()
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := &entity.User{
		ID:        uuid,
		Username:  username,
		Email:     email,
		Password:  string(hashedPass),
		IsActive:  true,
		Role:      "superuser",
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		CreatedBy: uuid,
		UpdatedBy: uuid,
	}
	err = userRepo.Create(ctx, user)
	if err != nil {
		panic(err)
	}
}
