package interfaces

import "github.com/pensk/invoices-api/internal/application/command"

type UserService interface {
	Authenticate(*command.AuthenticateUserCommand) (*command.AuthenticateUserResult, error)
}
