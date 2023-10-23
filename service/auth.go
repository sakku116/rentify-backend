package service

import (
	"context"
	"rentify/config"
	"rentify/entity"
	"rentify/exception"
	"rentify/helper"
	"rentify/repository"
)

type AuthService struct {
	userRepo repository.UserRepo
}

func NewAuthService(userRepo repository.UserRepo) AuthService {
	return AuthService{
		userRepo: userRepo,
	}
}

func (slf *AuthService) Login(username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", exception.AuthUsernameRequired
	}

	// check existance by username
	oldUser, err := slf.userRepo.GetByUsername(context.Background(), username)
	if err == exception.DbObjNotFound {
		return "", err
	}

	// check password
	isPwMatch := helper.ComparePasswordHash(password, oldUser.Password)
	if !isPwMatch {
		return "", exception.AuthPasswordIncorrect
	}

	// generate token
	newSessionID := helper.GenerateUUID()
	token, err := helper.GenerateJwtToken(username, newSessionID, config.Envs.JWT_SECRET, config.Envs.JWT_EXP)
	if err != nil {
		return "", err
	}

	// update session id from user
	err = slf.userRepo.Patch(context.Background(), oldUser.ID, &entity.User{
		SessionID: newSessionID,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (slf *AuthService) CheckToken() {

}
