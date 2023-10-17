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

func (self *AuthRepo) GenerateAccessToken(username string, session_id string) (string, error) {
	secretKey := []byte(config.Envs.JWT_SECRET)

	claims := jwt.MapClaims{
		"username":   "username",
		"exp":        time.Now().Add(time.Hour * time.Duration(config.Envs.JWT_EXP)).Unix(),
		"session_id": session_id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
