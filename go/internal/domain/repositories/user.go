package repositories

import "github.com/pensk/invoices-api/internal/infra/model"

type UserRepository interface {
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
}
