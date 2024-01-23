package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"github.com/pensk/invoices-api/internal/application/command"
	"github.com/pensk/invoices-api/internal/infra/model"
	"github.com/pensk/invoices-api/internal/interface/api/handler"
)

func TestInvoiceHandler_Create(t *testing.T) {
	t.Parallel()

	successBody := map[string]interface{}{
		"client_id":      1,
		"issue_date":     "2024-01-01",
		"due_date":       "2024-01-30",
		"payment_amount": 12000,
	}

	testCases := []struct {
		name           string
		reqBody        map[string]interface{}
		expectedStatus int
		runMock        func(*MockInvoiceService)
	}{
		{
			name:           "success",
			reqBody:        successBody,
			expectedStatus: http.StatusOK,
			runMock: func(is *MockInvoiceService) {
				issue, _ := time.Parse("2006-01-02", "2024-01-01")
				due, _ := time.Parse("2006-01-02", "2024-01-30")
				successCmd := &command.CreateInvoiceCommand{
					CompanyID:     1,
					ClientID:      1,
					IssueDate:     issue,
					DueDate:       due,
					PaymentAmount: 12000,
				}
				mockResponse := &command.CreateInvoiceResult{Invoice: &model.Invoice{ID: 1}}
				is.On("Create", successCmd).Return(mockResponse, nil).Once()
			},
		},
		{
			name:           "bad request",
			reqBody:        map[string]interface{}{"client_id": ""},
			expectedStatus: http.StatusBadRequest,
			runMock:        func(is *MockInvoiceService) {},
		},
	}

	for _, tc := range testCases {
		invoiceService := new(MockInvoiceService)
		router := chi.NewRouter()
		logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
		handler.NewInvoiceHandler(router, invoiceService, logger)

		tc.runMock(invoiceService)

		reqBodyBytes, _ := json.Marshal(tc.reqBody)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/", bytes.NewReader(reqBodyBytes))
		r.Header.Set("Content-Type", "application/json")

		ctx := context.WithValue(r.Context(), "company_id", 1)
		r = r.WithContext(ctx)
		ctx = context.WithValue(ctx, "user_id", 1)
		r = r.WithContext(ctx)

		router.ServeHTTP(w, r)

		if w.Code != tc.expectedStatus {
			t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Code)
		}

		invoiceService.AssertExpectations(t)
	}
}

func TestInvoiceHandler_List(t *testing.T) {
	t.Parallel()

	successBody := map[string]interface{}{
		"start_date": "2024-01-01",
		"end_date":   "2024-01-30",
		"page":       1,
		"per_page":   10,
	}
	successNoPage := map[string]interface{}{
		"start_date": "2024-01-01",
		"end_date":   "2024-01-30",
	}

	testCases := []struct {
		name           string
		reqBody        map[string]interface{}
		expectedStatus int
		runMock        func(*MockInvoiceService)
	}{
		{
			name:           "success",
			reqBody:        successBody,
			expectedStatus: http.StatusOK,
			runMock: func(is *MockInvoiceService) {
				startDate, _ := time.Parse("2006-01-02", "2024-01-01")
				endDate, _ := time.Parse("2006-01-02", "2024-01-30")
				successCmd := &command.ListInvoiceCommand{
					CompanyID: 1,
					StartDate: startDate,
					EndDate:   endDate,
					Page:      1,
					PerPage:   10,
				}
				mockResponse := &command.ListInvoiceResult{
					Invoices: []*model.Invoice{
						{
							ID:            1,
							CompanyID:     1,
							ClientID:      1,
							IssueDate:     "2024-01-01",
							PaymentAmount: 12000,
							FeeAmount:     480,
							FeeRate:       0.04,
							TaxAmount:     48,
							TaxRate:       0.10,
							TotalAmount:   12528,
							DueDate:       "2024-01-30",
							Status:        "pending",
						},
					},
					Count: 1,
				}
				is.On("List", successCmd).Return(mockResponse, nil).Once()
			},
		},
		{
			name:           "success default page",
			reqBody:        successNoPage,
			expectedStatus: http.StatusOK,
			runMock: func(is *MockInvoiceService) {
				startDate, _ := time.Parse("2006-01-02", "2024-01-01")
				endDate, _ := time.Parse("2006-01-02", "2024-01-30")
				successCmd := &command.ListInvoiceCommand{
					CompanyID: 1,
					StartDate: startDate,
					EndDate:   endDate,
					Page:      1,
					PerPage:   10,
				}
				mockResponse := &command.ListInvoiceResult{
					Invoices: []*model.Invoice{
						{
							ID:            1,
							CompanyID:     1,
							ClientID:      1,
							IssueDate:     "2024-01-01",
							PaymentAmount: 12000,
							FeeAmount:     480,
							FeeRate:       0.04,
							TaxAmount:     48,
							TaxRate:       0.10,
							TotalAmount:   12528,
							DueDate:       "2024-01-30",
							Status:        "pending",
						},
					},
					Count: 1,
				}
				is.On("List", successCmd).Return(mockResponse, nil).Once()
			},
		},
		{
			name:           "bad request",
			reqBody:        map[string]interface{}{"client_id": ""},
			expectedStatus: http.StatusBadRequest,
			runMock:        func(is *MockInvoiceService) {},
		},
	}

	for _, tc := range testCases {
		invoiceService := new(MockInvoiceService)
		router := chi.NewRouter()
		logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{}))
		handler.NewInvoiceHandler(router, invoiceService, logger)

		tc.runMock(invoiceService)

		reqBodyBytes, _ := json.Marshal(tc.reqBody)

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", bytes.NewReader(reqBodyBytes))
		r.Header.Set("Content-Type", "application/json")

		ctx := context.WithValue(r.Context(), "company_id", 1)
		r = r.WithContext(ctx)
		ctx = context.WithValue(ctx, "user_id", 1)
		r = r.WithContext(ctx)

		router.ServeHTTP(w, r)

		if w.Code != tc.expectedStatus {
			t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Code)
		}

		invoiceService.AssertExpectations(t)
	}
}
