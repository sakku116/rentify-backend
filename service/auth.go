package service

import (
	"context"
	"rentify/entity"
	"rentify/exception"
	"rentify/helper"
	"rentify/repository"
)

type AuthService struct {
	userRepo repository.UserRepo
	authRepo repository.AuthRepo
}

// Login authenticates a user with the given username and password.
//
// Parameters:
// - username: The username of the user.
// - password: The password of the user.
//
// Returns:
// - string: The authentication token if the login is successful.
// - error: An error if there is any issue during the login process.
//   - AuthUsernameRequired: If the username is empty.
//   - AuthPasswordRequired: If the password is empty.
//   - DbObjNotFound: If the user is not found in the database.
//   - AuthPasswordIncorrect: If the provided password is incorrect.
//   - InvalidToken: If the authentication token is invalid.
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
	token, err := slf.authRepo.GenerateAccessToken(username, newSessionID)

	// update session id
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
