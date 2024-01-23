package db_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/pensk/invoices-api/internal/infra/db"
	"github.com/pensk/invoices-api/internal/infra/model"
)

func TestUserRepository_GetByEmail(t *testing.T) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqldb.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "password", "company_id"}).
		AddRow(1, "test@example.com", "hashpassword", 3)

	mock.ExpectQuery(`SELECT id, email, password, company_id FROM users WHERE email = \?`).
		WithArgs("test@example.com").
		WillReturnRows(rows)

	repo := db.NewUserRepository(sqldb)
	user, err := repo.GetByEmail("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "hashpassword", user.PasswordHash)
	assert.Equal(t, 3, user.CompanyID)
}

func TestUserRepository_Create(t *testing.T) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqldb.Close()

	mock.ExpectExec(`INSERT INTO users \(company_id, name, email, password\) VALUES \(\?, \?, \?, \?\)`).
		WithArgs(3, "user name", "test@example.com", "hashedpassword").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := db.NewUserRepository(sqldb)
	user := &model.User{
		Name:         "user name",
		Email:        "test@example.com",
		PasswordHash: "hashedpassword",
		CompanyID:    3,
	}
	err = repo.Create(user)
	assert.Equal(t, 1, user.ID)
	assert.NoError(t, err)
}
