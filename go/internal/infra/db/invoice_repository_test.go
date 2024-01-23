package db_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pensk/invoices-api/internal/infra/db"
	"github.com/pensk/invoices-api/internal/infra/model"
	"github.com/stretchr/testify/assert"
)

func TestInvoiceRepository_Create(t *testing.T) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqldb.Close()

	invoice := &model.Invoice{
		CompanyID:     1,
		ClientID:      1,
		IssueDate:     time.Now().Format("2006-01-02"),
		PaymentAmount: 1000,
		FeeAmount:     40,
		FeeRate:       0.04,
		TaxAmount:     10,
		TaxRate:       0.1,
		TotalAmount:   1140,
		DueDate:       time.Now().AddDate(0, 0, 7).Format("2006-01-02"),
		Status:        string(model.InvoiceStatusPending),
	}

	mock.ExpectExec("^INSERT INTO invoices (.+)$").WithArgs(
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
	).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := db.NewInvoiceRepository(sqldb)
	err = repo.Create(invoice)

	assert.NoError(t, err)
	mock.ExpectationsWereMet()
}
