package services

import (
	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/pensk/invoices-api/internal/domain/repositories"
)

type InvoiceService struct {
	ir repositories.InvoiceRepository
}

func NewInvoiceService(ir repositories.InvoiceRepository) *InvoiceService {
	return &InvoiceService{ir: ir}
}

func (s *InvoiceService) Create(*command.CreateInvoiceCommand) (*command.CreateInvoiceResult, error) {
	return nil, nil
}

func (s *InvoiceService) List(*command.ListInvoiceCommand) (*command.ListInvoiceResult, error) {
	return nil, nil
}
