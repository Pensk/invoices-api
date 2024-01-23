package command

import (
	"time"

	"github.com/pensk/invoices-api/internal/infra/model"
)

type CreateInvoiceCommand struct {
	CompanyID     int
	ClientID      int
	IssueDate     time.Time
	PaymentAmount uint64
	DueDate       time.Time
	Status        string
}

type CreateInvoiceResult struct {
	Invoice *model.Invoice `json:"invoice"`
}

type ListInvoiceCommand struct {
	CompanyID int
	StartDate time.Time
	EndDate   time.Time
	PerPage   int
	Page      int
}

type ListInvoiceResult struct {
	Invoices []*model.Invoice `json:"invoices"`
	Count    int              `json:"count"`
}
