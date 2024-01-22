package model

type User struct {
	ID           int
	CompanyID    int
	PasswordHash string
	Name         string
	Email        string
}
