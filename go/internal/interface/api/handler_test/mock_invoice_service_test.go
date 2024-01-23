package handler_test

import (
	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/stretchr/testify/mock"
)

type MockInvoiceService struct {
	mock.Mock
}

func (m *MockInvoiceService) Create(cmd *command.CreateInvoiceCommand) (*command.CreateInvoiceResult, error) {
	args := m.Called(cmd)
	return args.Get(0).(*command.CreateInvoiceResult), args.Error(1)
}

func (m *MockInvoiceService) List(cmd *command.ListInvoiceCommand) (*command.ListInvoiceResult, error) {
	args := m.Called(cmd)
	return args.Get(0).(*command.ListInvoiceResult), args.Error(1)
}
