package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/pensk/invoices-api/internal/application/interfaces"
	"github.com/pensk/invoices-api/internal/domain/repositories"
	"github.com/pensk/invoices-api/internal/infra/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ur repositories.UserRepository
	cr repositories.CompanyRepository
}

func NewUserService(ur repositories.UserRepository, cr repositories.CompanyRepository) interfaces.UserService {
	return &UserService{
		ur: ur,
		cr: cr,
	}
}

const (
	ExpirationHours = 74
)

func generateToken(userID int) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * ExpirationHours).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}

func (s *UserService) Authenticate(cmd *command.AuthenticateUserCommand) (*command.AuthenticateUserResult, error) {
	user, err := s.ur.GetByEmail(cmd.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(cmd.Password))
	if err != nil {
		return nil, err
	}

	token, err := generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &command.AuthenticateUserResult{AccessToken: token}, nil
}

func (s *UserService) Create(cmd *command.CreateUserCommand) (*command.CreateUserResult, error) {
	company, err := s.cr.GetByID(cmd.CompanyID)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:         cmd.Name,
		Email:        cmd.Email,
		PasswordHash: string(hashedPassword),
		CompanyID:    company.ID,
	}

	err = s.ur.Create(user)
	if err != nil {
		return nil, err
	}

	token, err := generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &command.CreateUserResult{AccessToken: token}, nil
}
