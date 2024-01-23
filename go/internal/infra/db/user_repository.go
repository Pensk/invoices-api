package db

import (
	"database/sql"

	"github.com/pensk/invoices-api/internal/domain/repositories"
	"github.com/pensk/invoices-api/internal/infra/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, email, password_hash, company_id FROM users WHERE email = ?`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CompanyID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	stmt := `INSERT INTO users (company_id, name, email, password_hash) VALUES (?, ?, ?, ?)`

	result, err := r.db.Exec(stmt, user.CompanyID, user.Name, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}
