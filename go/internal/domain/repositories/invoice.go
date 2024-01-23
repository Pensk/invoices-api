package repositories

import "github.com/pensk/invoices-api/internal/infra/model"

type InvoiceRepository interface {
	Create(invoice *model.Invoice) error
	List() error
}
