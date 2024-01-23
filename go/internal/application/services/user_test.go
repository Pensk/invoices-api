package services_test

import (
	"testing"

	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/pensk/invoices-api/internal/application/services"
	"github.com/pensk/invoices-api/internal/infra/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	user.ID = 1
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id int) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

type MockCompanyRepository struct {
	mock.Mock
}

func (m *MockCompanyRepository) GetByID(id int) (*model.Company, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Company), args.Error(1)
}

func TestUserService_Create(t *testing.T) {
	userRepo := new(MockUserRepository)
	companyRepo := new(MockCompanyRepository)
	userService := services.NewUserService(userRepo, companyRepo)

	cmd := &command.CreateUserCommand{
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "password",
		CompanyID: 1,
	}

	company := &model.Company{ID: 1}

	companyRepo.On("GetByID", company.ID).Return(company, nil)
	userRepo.On("Create", mock.AnythingOfType("*model.User")).Return(nil)

	res, err := userService.Create(cmd)
	assert.NoError(t, err)

	if res.AccessToken == "" {
		t.Errorf("Expected access token to be generated")
	}

	companyRepo.AssertExpectations(t)
	userRepo.AssertExpectations(t)
}

func TestUserService_Authenticate(t *testing.T) {
	userRepo := new(MockUserRepository)
	companyRepo := new(MockCompanyRepository)
	userService := services.NewUserService(userRepo, companyRepo)

	cmd := &command.AuthenticateUserCommand{
		Email:    "test@example.com",
		Password: "password",
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)

	user := &model.User{ID: 1, PasswordHash: string(hashedPass)}

	userRepo.On("GetByEmail", cmd.Email).Return(user, nil)

	res, err := userService.Authenticate(cmd)
	assert.NoError(t, err)

	if res.AccessToken == "" {
		t.Errorf("Expected access token to be generated")
	}

	companyRepo.AssertExpectations(t)
	userRepo.AssertExpectations(t)
}
