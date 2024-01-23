package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/pensk/invoices-api/internal/interface/api/handler"
)

func TestAuthenticateUser(t *testing.T) {
	t.Parallel()

	successBody := map[string]interface{}{"Email": "test@ex.com", "Password": "password"}
	unauthBody := map[string]interface{}{"Email": "text@ex.com", "Password": "unset"}

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
			handler.NewUserHandler(router, userService)

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
