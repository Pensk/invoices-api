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
	return nil, nil
}
