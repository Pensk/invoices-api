package services

import (
	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/pensk/invoices-api/internal/application/interfaces"
	"github.com/pensk/invoices-api/internal/domain/repositories"
)

type UserService struct {
	ur repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) interfaces.UserService {
	return &UserService{
		ur: ur,
	}
}

func (s *UserService) Authenticate(cmd *command.AuthenticateUserCommand) (*command.AuthenticateUserResult, error) {
	_, err := s.ur.GetByEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	return &command.AuthenticateUserResult{}, nil
}
