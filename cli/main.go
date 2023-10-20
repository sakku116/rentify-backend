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

func SeedSuperuser(userRepo *repository.UserRepo) {
	ctx := context.Background()

	// check for existing superuser
	existingUser, err := userRepo.GetByUsername(ctx, "superuser")
	if err != nil && err != exception.DbObjNotFound {
		panic(err)
	}
	if existingUser != nil {
		fmt.Println("superuser already exists")
		return
	}

	uuid := helper.GenerateUUID()
	timeNow := time.Now().Unix()
	hashedPass, err := bcrypt.GenerateFromPassword([]byte("superuser"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := &entity.User{
		ID:        uuid,
		Username:  "superuser",
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
