package service

import (
	"context"
	"rentify/exception"
	"rentify/helper"
	"rentify/repository"
)

type AuthService struct {
	userRepo *repository.UserRepo
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
func (self *AuthService) Login(username string, password string) (string, error) {
	if username == "" || password == "" {
		err := exception.AuthUsernameRequired
		return "", err
	}

	user, err := self.userRepo.GetByUsername(context.Background(), username)
	if err == exception.DbObjNotFound {
		return "", err
	}

	isPwMatch := helper.ComparePasswordHash(password, user.Password)
	if !isPwMatch {
		return "", exception.AuthPasswordIncorrect
	}

	return "", nil
}

func (self *AuthService) CheckToken() {

}
