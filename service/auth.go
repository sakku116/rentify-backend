package service

import (
	"context"
	"rentify/config"
	"rentify/entity"
	"rentify/exception"
	"rentify/helper"
	"rentify/repository"

	"gopkg.in/mgo.v2/bson"
)

type AuthService struct {
	userRepo repository.UserRepo
}

func NewAuthService(userRepo repository.UserRepo) AuthService {
	return AuthService{
		userRepo: userRepo,
	}
}

func (slf *AuthService) Login(ctx context.Context, username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", exception.AuthUserPassRequired
	}

	// check existance by username
	oldUser, err := slf.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// check password
	isPwMatch := helper.ComparePasswordHash(password, oldUser.Password)
	if !isPwMatch {
		return "", exception.AuthPasswordIncorrect
	}

	// generate token
	newSessionID := helper.GenerateUUID()
	token, err := helper.GenerateJwtToken(username, oldUser.ID, oldUser.Role, newSessionID, config.Envs.JWT_SECRET, config.Envs.JWT_EXP)
	if err != nil {
		return "", err
	}

	// update session id from user
	err = slf.userRepo.Patch(ctx, oldUser.ID, bson.M{
		"session_id": newSessionID,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

/*
raises:
- exception.AuthInvalidToken
- exception.AuthUserNotFound
- exception.AuthUserBanned
*/
func (slf *AuthService) CheckToken(ctx context.Context, token string) (*entity.User, error) {
	claims, err := helper.ValidateJWT(token)
	if err != nil {
		return nil, exception.AuthInvalidToken
	}

	userID := claims["user_id"].(string)
	user, err := slf.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, exception.AuthUserNotFound
	}

	if user.IsActive == false {
		return nil, exception.AuthUserBanned
	}

	return user, nil
}

/*
raises:
- exception.AuthInvalidRole
- exception.DBObjNotFound
*/
func (slf *AuthService) SetRole(ctx context.Context, user_id string, role string) error {
	// validate role
	if role != "owner" && role != "customer" {
		return exception.AuthInvalidRole
	}

	// update user role
	err := slf.userRepo.Patch(ctx, user_id, bson.M{
		role: "superuser",
	})
	if err != nil {
		return err
	}

	return nil
}
