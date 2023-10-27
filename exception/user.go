package exception

import "fmt"

var (
	UserAlreadyExistByUsername = fmt.Errorf("user already exists")
	UserAlreadyExistByEmail    = fmt.Errorf("user already exists")
)
