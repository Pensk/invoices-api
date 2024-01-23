package request

import (
	"errors"
	"time"

	"github.com/pensk/invoices-api/internal/application/command"
)

type CreateInvoiceRequest struct {
	ClientID      int    `json:"client_id"`
	IssueDate     string `json:"issue_date"`
	PaymentAmount uint64 `json:"payment_amount"`
	DueDate       string `json:"due_date"`
}

func (r *CreateInvoiceRequest) ToCreateInvoiceCommand() (*command.CreateInvoiceCommand, error) {
	if r.ClientID == 0 || r.IssueDate == "" || r.PaymentAmount == 0 || r.DueDate == "" {
		return nil, errors.New("invalid request")
	}

	issueDate, err := time.Parse("2006-01-02", r.IssueDate)
	if err != nil {
		return nil, err
	}
	dueDate, err := time.Parse("2006-01-02", r.IssueDate)
	if err != nil {
		return nil, err
	}

	cmd := &command.CreateInvoiceCommand{
		ClientID:      r.ClientID,
		IssueDate:     issueDate,
		PaymentAmount: r.PaymentAmount,
		DueDate:       dueDate,
	}

	return cmd, nil
}
