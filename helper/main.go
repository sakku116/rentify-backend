package helper

import (
	"fmt"
	"reflect"
	"rentify/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUUID() string {
	return uuid.New().String()
}

func ParseBsonPatchStruct(obj interface{}) map[string]interface{} {
	var result = make(map[string]interface{})
	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		tag := field.Tag.Get("bson_patch")

		if tag != "" {
			tagParts := strings.Split(tag, ",")
			tagName := tagParts[0]
			omitempty := false

			if len(tagParts) > 1 && tagParts[1] == "omitempty" {
				omitempty = true
			}

			if omitempty && IsZero(val.Field(i)) {
				continue
			}

			value := val.Field(i).Interface()
			result[tagName] = value
		}
	}

	return result
}

func IsZero(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	case reflect.Slice, reflect.Map, reflect.String:
		return value.Len() == 0
	default:
		zeroValue := reflect.Zero(value.Type()).Interface()
		return reflect.DeepEqual(value.Interface(), zeroValue)
	}
}

/*
return example

	{
		"username": "fulan",
		"user_id": "1234",
		"session_id": "1234",
		"exp": 12345,
	}
*/
func GenerateJwtToken(username string, user_id string, session_id string, secret_key string, exp int) (string, error) {
	secretKey := []byte(secret_key)

	claims := jwt.MapClaims{
		"username":   username,
		"user_id":    user_id,
		"exp":        time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
		"session_id": session_id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SIGNATURE_KEY = []byte(config.Envs.JWT_SECRET)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
