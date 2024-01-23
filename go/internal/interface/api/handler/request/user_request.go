package request

import (
	"errors"

	"github.com/pensk/invoices-api/internal/application/command"
)

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthenticateUserRequest) ToAuthenticateUserCommand() (*command.AuthenticateUserCommand, error) {
	if r.Email == "" || r.Password == "" {
		return nil, errors.New("invalid request")
	}

	cmd := &command.AuthenticateUserCommand{
		Email:    r.Email,
		Password: r.Password,
	}

	return cmd, nil
}

type CreateUserRequest struct {
	CompanyID int    `json:"company_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (r *CreateUserRequest) ToCreateUserCommand() (*command.CreateUserCommand, error) {
	if r.Name == "" || r.Email == "" || r.Password == "" {
		return nil, errors.New("invalid request")
	}

	cmd := &command.CreateUserCommand{
		CompanyID: r.CompanyID,
		Name:      r.Name,
		Email:     r.Email,
		Password:  r.Password,
	}

	return cmd, nil
}
