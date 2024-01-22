package request

import "github.com/pensk/invoices-api/internal/application/command"

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthenticateUserRequest) ToAuthenticateUserCommand() *command.AuthenticateUserCommand {
	cmd := &command.AuthenticateUserCommand{
		Email:    r.Email,
		Password: r.Password,
	}

	return cmd
}
