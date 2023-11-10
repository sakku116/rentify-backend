package service

import (
	"context"
	"fmt"
	"rentify/config"
	"rentify/domain/entity"
	"rentify/repository"
	error_utils "rentify/utils/error"
	"rentify/utils/helper"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type AuthService struct {
	userRepo repository.IUserRepo
}

type IAuthService interface {
	Login(ctx context.Context, username string, password string) (string, error)
	CheckToken(ctx context.Context, token string) (*entity.User, error)
	SetRole(ctx context.Context, user_id string, role string) error
	Register(ctx context.Context, username string, email string, password string) error
}

func NewAuthService(userRepo repository.IUserRepo) IAuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (slf *AuthService) Login(ctx context.Context, username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", &error_utils.CustomErr{
			Code:    400,
			Message: "username and password are required",
		}
	}

	// check existance by username
	oldUser, err := slf.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", &error_utils.CustomErr{
			Code:    400,
			Message: "username not found",
		}
	}

	// check password
	isPwMatch := helper.ComparePasswordHash(password, oldUser.Password)
	if !isPwMatch {
		return "", &error_utils.CustomErr{
			Code:    401,
			Message: "password incorrect",
		}
	}

	// generate token
	newSessionID := helper.GenerateUUID()
	token, err := helper.GenerateJwtToken(username, oldUser.ID, oldUser.Role, newSessionID, config.Envs.JWT_SECRET, config.Envs.JWT_EXP)
	if err != nil {
		return "", &error_utils.CustomErr{
			Code:    500,
			Message: "error when generating jwt token",
		}
	}

	// update session id from user
	err = slf.userRepo.Patch(ctx, oldUser.ID, bson.M{
		"session_id": newSessionID,
	})
	if err != nil {
		return "", &error_utils.CustomErr{
			Code:    404,
			Message: fmt.Sprintf("user with id %s not found", oldUser.ID),
		}
	}

	return token, nil
}

func (slf *AuthService) CheckToken(ctx context.Context, token string) (*entity.User, error) {
	claims, err := helper.ValidateJWT(token)
	if err != nil {
		return nil, &error_utils.CustomErr{
			Code:    401,
			Message: "invalid token",
		}
	}

	userID := claims["user_id"].(string)
	user, err := slf.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, &error_utils.CustomErr{
			Code:    404,
			Message: "user not found",
		}
	}

	if !user.IsActive {
		return nil, &error_utils.CustomErr{
			Code:    403,
			Message: "user is banned",
		}
	}

	if user.Role == "" {
		return nil, &error_utils.CustomErr{
			Code:    403,
			Message: "role is not set",
		}
	}

	return user, nil
}

func (slf *AuthService) SetRole(ctx context.Context, user_id string, role string) error {
	// validate role
	if role != "owner" && role != "customer" {
		return &error_utils.CustomErr{
			Code:    400,
			Message: "invalid role",
		}
	}

	// update user role
	err := slf.userRepo.Patch(ctx, user_id, bson.M{
		role: "superuser",
	})
	if err != nil {
		return &error_utils.CustomErr{
			Code:    404,
			Message: fmt.Sprintf("user with id %s not found", user_id),
		}
	}

	return nil
}

func (slf *AuthService) Register(ctx context.Context, username string, email string, password string) error {
	// user & email existance validation
	userByUsername, _ := slf.userRepo.GetByUsername(ctx, username)
	if userByUsername != nil {
		return &error_utils.CustomErr{
			Code:    400,
			Message: "username already exist",
		}
	}
	userByEmail, _ := slf.userRepo.GetByEmail(ctx, email)
	if userByEmail != nil {
		return &error_utils.CustomErr{
			Code:    400,
			Message: "email already exist",
		}
	}

	// hash password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// new object creation
	uuid := helper.GenerateUUID()
	timeNow := time.Now().Unix()
	newUser := &entity.User{
		ID:        uuid,
		Username:  username,
		Email:     email,
		Password:  string(hashedPass),
		IsActive:  true,
		Role:      "customer",
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		CreatedBy: uuid,
		UpdatedBy: uuid,
	}
	err = slf.userRepo.Create(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}
