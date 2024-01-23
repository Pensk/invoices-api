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
	company := &model.Company{}
	stmt := `SELECT id, name FROM companies WHERE id = ?`

	err := r.db.QueryRow(stmt, id).Scan(&company.ID, &company.Name)
	if err != nil {
		return nil, err
	}
	return company, nil
}
