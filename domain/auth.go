package domain

type AuthService interface {
	GenerateToken() (string, error)
}
