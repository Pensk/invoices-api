package db

import (
	"database/sql"

	"github.com/pensk/invoices-api/internal/infra/model"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Create(invoice *model.Invoice) error {
	return nil
}

func (r *InvoiceRepository) List() error {
	return nil
}
