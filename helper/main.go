package helper

import (
	"reflect"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUUID() string {
	return uuid.New().String()
}

func ParseBsonPatchStruct(obj interface{}) map[string]interface{} {
	var result map[string]interface{}
	val := reflect.ValueOf(obj)

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
			if omitempty == false {
				result[tagName] = value
			}
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
