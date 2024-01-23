package db

import (
	"database/sql"

	"github.com/pensk/invoices-api/internal/domain/repositories"
	"github.com/pensk/invoices-api/internal/infra/model"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) repositories.CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (r *CompanyRepository) GetByID(id int) (*model.Company, error) {
	return nil, nil
}
