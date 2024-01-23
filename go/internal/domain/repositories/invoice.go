package repositories

import "github.com/pensk/invoices-api/internal/infra/model"

type InvoiceRepository interface {
	Create(invoice *model.Invoice) error
	List(companyID int, startDate string, endDate string, page int, perPage int) ([]*model.Invoice, error)
	Count(companyID int, startDate string, endDate string) (int, error)
}
