package repository

import (
	"rentify/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthRepo struct {
}

func NewAuthRepo() *AuthRepo {
	return &AuthRepo{}
}

func (self *AuthRepo) GenerateAccessToken(username string, exp_hours int, token_id string) (string, error) {
	secretKey := []byte(config.Envs.JWT_SECRET)

	claims := jwt.MapClaims{
		"username": "username",
		"exp":      time.Now().Add(time.Hour * time.Duration(exp_hours)).Unix(),
		"token_id": token_id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
