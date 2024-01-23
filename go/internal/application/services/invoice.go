package services

import (
	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/pensk/invoices-api/internal/domain/repositories"
	"github.com/pensk/invoices-api/internal/infra/model"
)

const (
	FeeRate = 0.04
	TaxRate = 0.10
)

type InvoiceService struct {
	ir repositories.InvoiceRepository
}

func NewInvoiceService(ir repositories.InvoiceRepository) *InvoiceService {
	return &InvoiceService{ir: ir}
}

func (s *InvoiceService) Create(cmd *command.CreateInvoiceCommand) (*command.CreateInvoiceResult, error) {
	feeAmount := float64(cmd.PaymentAmount) * FeeRate
	taxAmount := feeAmount * TaxRate
	totalAmount := float64(cmd.PaymentAmount) + feeAmount + taxAmount
	status := model.InvoiceStatusPending
	issueDate := cmd.IssueDate.Format("2006-01-02")
	dueDate := cmd.DueDate.Format("2006-01-02")

	invoice := &model.Invoice{
		CompanyID:     cmd.CompanyID,
		ClientID:      cmd.ClientID,
		IssueDate:     issueDate,
		PaymentAmount: cmd.PaymentAmount,
		FeeAmount:     uint64(feeAmount),
		FeeRate:       FeeRate,
		TaxAmount:     uint64(taxAmount),
		TaxRate:       TaxRate,
		TotalAmount:   uint64(totalAmount),
		DueDate:       dueDate,
		Status:        string(status),
	}

	if err := s.ir.Create(invoice); err != nil {
		return nil, err
	}

	return &command.CreateInvoiceResult{Invoice: invoice}, nil
}

func (s *InvoiceService) List(cmd *command.ListInvoiceCommand) (*command.ListInvoiceResult, error) {
	startDate := cmd.StartDate.Format("2006-01-02")
	endDate := cmd.EndDate.Format("2006-01-02")

	count, err := s.ir.Count(cmd.CompanyID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	invoices, err := s.ir.List(cmd.CompanyID, startDate, endDate, cmd.Page, cmd.PerPage)
	if err != nil {
		return nil, err
	}

	return &command.ListInvoiceResult{Invoices: invoices, Count: count}, nil
}
