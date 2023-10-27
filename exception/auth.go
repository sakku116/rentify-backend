package exception

import "fmt"

var (
	AuthUsernameRequired  = fmt.Errorf("username required")
	AuthPasswordRequired  = fmt.Errorf("password required")
	AuthUsernameNotFound  = fmt.Errorf("username not found")
	AuthUserNotFound      = fmt.Errorf("user not found")
	AuthPasswordIncorrect = fmt.Errorf("password incorrect")
	AuthInvalidToken      = fmt.Errorf("invalid token")
	AuthUserPassRequired  = fmt.Errorf("username and password are required")
	AuthUserBanned        = fmt.Errorf("user is banned")
	AuthRoleRequired      = fmt.Errorf("role is not set")
	AuthInvalidRole       = fmt.Errorf("invalid role")
)
