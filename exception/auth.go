package exception

import "fmt"

var (
	AuthUsernameRequired  = fmt.Errorf("username required")
	AuthPasswordRequired  = fmt.Errorf("password required")
	AuthUsernameNotFound  = fmt.Errorf("username not found")
	AuthPasswordIncorrect = fmt.Errorf("password incorrect")
	AuthInvalidToken      = fmt.Errorf("invalid token")
)
