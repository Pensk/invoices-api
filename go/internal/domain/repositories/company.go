package repositories

import "github.com/pensk/invoices-api/internal/infra/model"

type CompanyRepository interface {
	GetByID(id int) (*model.Company, error)
}
