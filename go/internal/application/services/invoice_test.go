package services_test

import (
	"testing"
	"time"

	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/pensk/invoices-api/internal/application/services"
	"github.com/pensk/invoices-api/internal/infra/model"
	"github.com/stretchr/testify/assert"
)

type MockInvoiceRepository struct {
	Invoice      *model.Invoice
	Invoices     []*model.Invoice
	InvoiceCount int
}

func (m *MockInvoiceRepository) Create(invoice *model.Invoice) error {
	m.Invoice = invoice
	return nil
}

func (m *MockInvoiceRepository) List(companyID int, startDate string, endDate string, page int, perPage int) ([]*model.Invoice, error) {
	return m.Invoices, nil
}

func (m *MockInvoiceRepository) Count(companyID int, startDate string, endDate string) (int, error) {
	return m.InvoiceCount, nil
}

func TestInvoiceService_Create(t *testing.T) {
	invoiceRepo := new(MockInvoiceRepository)
	service := services.NewInvoiceService(invoiceRepo)

	cmd := &command.CreateInvoiceCommand{
		CompanyID:     1,
		ClientID:      1,
		IssueDate:     time.Now(),
		PaymentAmount: 1000,
		DueDate:       time.Now().AddDate(0, 0, 7),
	}

	_, err := service.Create(cmd)

	floatAmt := float64(cmd.PaymentAmount)

	assert.NoError(t, err)
	assert.Equal(t, uint64(floatAmt*services.FeeRate), invoiceRepo.Invoice.FeeAmount)
	assert.Equal(t, uint64(float64(floatAmt*services.FeeRate)*services.TaxRate), invoiceRepo.Invoice.TaxAmount)
	assert.Equal(t, uint64(floatAmt+(float64(floatAmt*services.FeeRate)*(1+services.TaxRate))), invoiceRepo.Invoice.TotalAmount)
	assert.Equal(t, string(model.InvoiceStatusPending), invoiceRepo.Invoice.Status)
}

func TestInvoiceService_List(t *testing.T) {
	mockRepo := &MockInvoiceRepository{
		Invoices: []*model.Invoice{
			{ID: 1, CompanyID: 1, ClientID: 1, IssueDate: time.Now().Format("2006-01-02"), PaymentAmount: 1000, DueDate: time.Now().AddDate(0, 0, 7).Format("2006-01-02"), Status: string(model.InvoiceStatusPending)},
			{ID: 2, CompanyID: 1, ClientID: 2, IssueDate: time.Now().Format("2006-01-02"), PaymentAmount: 2000, DueDate: time.Now().AddDate(0, 0, 7).Format("2006-01-02"), Status: string(model.InvoiceStatusPending)},
		},
		InvoiceCount: 2,
	}
	service := services.NewInvoiceService(mockRepo)

	cmd := &command.ListInvoiceCommand{
		CompanyID: 1,
		StartDate: time.Now().AddDate(-1, 0, 0),
		EndDate:   time.Now(),
		PerPage:   10,
		Page:      1,
	}

	result, err := service.List(cmd)

	assert.NoError(t, err)
	assert.Equal(t, mockRepo.Invoices, result.Invoices)
	assert.Equal(t, mockRepo.InvoiceCount, result.Count)
}
