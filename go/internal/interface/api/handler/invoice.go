package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pensk/invoices-api/internal/application/interfaces"
	"github.com/pensk/invoices-api/internal/interface/api/handler/request"
)

type InvoiceHandler struct {
	service interfaces.InvoiceService
	logger  *slog.Logger
}

func NewInvoiceHandler(router chi.Router, service interfaces.InvoiceService, logger *slog.Logger) {
	handler := &InvoiceHandler{
		service: service,
		logger:  logger,
	}

	router.Post("/", handler.Create)
	router.Get("/", handler.List)
}

func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req *request.CreateInvoiceRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd, err := req.ToCreateInvoiceCommand()
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd.CompanyID = ctx.Value("company_id").(int)

	res, err := h.service.Create(cmd)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *InvoiceHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req *request.ListInvoiceRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd, err := req.ToListInvoiceCommand()
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd.CompanyID = ctx.Value("company_id").(int)

	res, err := h.service.List(cmd)
	if err != nil {
		h.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
