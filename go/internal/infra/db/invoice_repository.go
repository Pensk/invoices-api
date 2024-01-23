package db

import (
	"database/sql"
	"fmt"

	"github.com/pensk/invoices-api/internal/infra/model"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Create(invoice *model.Invoice) error {
	stmt := `INSERT INTO invoices (
		company_id, 
		client_id,
		issue_date,
		payment_amount,
		fee,
		fee_rate,
		tax,
		tax_rate,
		total_amount,
		due_date,
		status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(
		stmt,
		invoice.CompanyID,
		invoice.ClientID,
		invoice.IssueDate,
		invoice.PaymentAmount,
		invoice.FeeAmount,
		invoice.FeeRate,
		invoice.TaxAmount,
		invoice.TaxRate,
		invoice.TotalAmount,
		invoice.DueDate,
		invoice.Status,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	invoice.ID = int(id)
	return nil
}

func (r *InvoiceRepository) List(companyID int, startDate string, endDate string, page int, perPage int) ([]*model.Invoice, error) {
	invoices := []*model.Invoice{}
	stmt := `SELECT 
		id,
		issue_date,
		payment_amount,
		fee,
		fee_rate,
		tax,
		tax_rate,
		total_amount,
		due_date,
		status,
		company_id,
		client_id
	FROM invoices WHERE company_id = ? AND due_date >= ? AND due_date <= ? LIMIT ? OFFSET ?`

	fmt.Printf("Query: %s\n", stmt)
	fmt.Printf("Parameters: %d, %s, %s, %d, %d\n", companyID, startDate, endDate, perPage, (page-1)*perPage)

	rows, err := r.db.Query(stmt, companyID, startDate, endDate, perPage, (page-1)*perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		invoice := &model.Invoice{}
		err := rows.Scan(
			&invoice.ID,
			&invoice.IssueDate,
			&invoice.PaymentAmount,
			&invoice.FeeAmount,
			&invoice.FeeRate,
			&invoice.TaxAmount,
			&invoice.TaxRate,
			&invoice.TotalAmount,
			&invoice.DueDate,
			&invoice.Status,
			&invoice.CompanyID,
			&invoice.ClientID,
		)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (r *InvoiceRepository) Count(companyID int, startDate string, endDate string) (int, error) {
	stmt := `SELECT COUNT(*) FROM invoices WHERE company_id = ? AND due_date >= ? AND due_date <= ?`

	var count int
	err := r.db.QueryRow(stmt, companyID, startDate, endDate).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
