package interfaces

import "github.com/pensk/invoices-api/internal/application/command"

type InvoiceService interface {
	Create(*command.CreateInvoiceCommand) (*command.CreateInvoiceResult, error)
	List(*command.ListInvoiceCommand) (*command.ListInvoiceResult, error)
}
