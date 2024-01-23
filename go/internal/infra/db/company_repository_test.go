package db_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/pensk/invoices-api/internal/infra/db"
)

func TestCompanyRepository_GetByID(t *testing.T) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqldb.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Test Company")

	mock.ExpectQuery(`SELECT id, name FROM companies WHERE id = \?`).
		WithArgs(1).
		WillReturnRows(rows)

	repo := db.NewCompanyRepository(sqldb)
	company, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, company.ID)
	assert.Equal(t, "Test Company", company.Name)
}
