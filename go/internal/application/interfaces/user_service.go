package interfaces

import "github.com/pensk/invoices-api/internal/application/command"

type UserService interface {
	Create(*command.CreateUserCommand) (*command.CreateUserResult, error)
	Authenticate(*command.AuthenticateUserCommand) (*command.AuthenticateUserResult, error)
}
