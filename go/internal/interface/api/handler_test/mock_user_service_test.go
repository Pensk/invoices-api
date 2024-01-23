package handler_test

import (
	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Authenticate(cmd *command.AuthenticateUserCommand) (*command.AuthenticateUserResult, error) {
	args := m.Called(cmd)
	return args.Get(0).(*command.AuthenticateUserResult), args.Error(1)
}

func (m *MockUserService) Create(cmd *command.CreateUserCommand) (*command.CreateUserResult, error) {
	args := m.Called(cmd)
	return args.Get(0).(*command.CreateUserResult), args.Error(1)
}
