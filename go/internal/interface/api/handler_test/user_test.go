package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi"
	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/pensk/invoices-api/internal/interface/api/handler"
)

func TestAuthenticateUser(t *testing.T) {
	t.Parallel()

	successBody := map[string]interface{}{"email": "test@ex.com", "password": "password"}
	unauthBody := map[string]interface{}{"email": "text@ex.com", "password": "unset"}

	testCases := []struct {
		name           string
		reqBody        map[string]interface{}
		expectedStatus int
		runMock        func(*MockUserService)
	}{
		{
			name:           "success",
			reqBody:        successBody,
			expectedStatus: http.StatusOK,
			runMock: func(us *MockUserService) {
				successCmd := &command.AuthenticateUserCommand{Email: "test@ex.com", Password: "password"}
				mockResponse := &command.AuthenticateUserResult{AccessToken: "token"}
				us.On("Authenticate", successCmd).Return(mockResponse, nil).Once()
			},
		},
		{
			name:           "bad request",
			reqBody:        map[string]interface{}{"Username": ""},
			expectedStatus: http.StatusBadRequest,
			runMock:        func(us *MockUserService) {},
		},
		{
			name:           "unauthorized",
			reqBody:        unauthBody,
			expectedStatus: http.StatusUnauthorized,
			runMock: func(us *MockUserService) {
				mockError := errors.New("unauthorized")
				unauthCmd := &command.AuthenticateUserCommand{Email: "text@ex.com", Password: "unset"}
				us.On("Authenticate", unauthCmd).Return(&command.AuthenticateUserResult{}, mockError).Once()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userService := new(MockUserService)
			router := chi.NewRouter()
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
			handler.NewUserHandler(router, userService, logger)

			tc.runMock(userService)

			reqBodyBytes, _ := json.Marshal(tc.reqBody)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/users/authenticate", bytes.NewReader(reqBodyBytes))
			r.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, r)

			if w.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Code)
			}

			userService.AssertExpectations(t)
		})
	}
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	successBody := map[string]interface{}{"company_id": 1, "name": "new user", "email": "test@ex.com", "Password": "password"}
	failBody := map[string]interface{}{"company_id": 44, "name": "new user", "email": "text@ex.com", "Password": "unset"}

	testCases := []struct {
		name           string
		reqBody        map[string]interface{}
		expectedStatus int
		runMock        func(*MockUserService)
	}{
		{
			name:           "success",
			reqBody:        successBody,
			expectedStatus: http.StatusOK,
			runMock: func(us *MockUserService) {
				successCmd := &command.CreateUserCommand{CompanyID: 1, Name: "new user", Email: "test@ex.com", Password: "password"}
				mockResponse := &command.CreateUserResult{AccessToken: "token"}
				us.On("Create", successCmd).Return(mockResponse, nil).Once()
			},
		},
		{
			name:           "bad request",
			reqBody:        map[string]interface{}{"Username": ""},
			expectedStatus: http.StatusBadRequest,
			runMock:        func(us *MockUserService) {},
		},
		{
			name:           "invalid company id",
			reqBody:        failBody,
			expectedStatus: http.StatusBadRequest,
			runMock: func(us *MockUserService) {
				mockError := errors.New("invalid company id")
				unauthCmd := &command.CreateUserCommand{CompanyID: 44, Name: "new user", Email: "text@ex.com", Password: "unset"}
				us.On("Create", unauthCmd).Return(&command.CreateUserResult{}, mockError).Once()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			userService := new(MockUserService)
			router := chi.NewRouter()
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
			handler.NewUserHandler(router, userService, logger)

			tc.runMock(userService)

			reqBodyBytes, _ := json.Marshal(tc.reqBody)

			w := httptest.NewRecorder()
			r := httptest.NewRequest("POST", "/users/signup", bytes.NewReader(reqBodyBytes))
			r.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, r)

			if w.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Code)
			}

			userService.AssertExpectations(t)
		})
	}
}
