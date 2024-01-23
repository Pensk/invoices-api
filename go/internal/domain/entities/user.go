package entities

type User struct {
	ID           int
	CompanyID    int
	Name         string
	Email        string
	PasswordHash string
}
