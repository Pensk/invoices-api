package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pensk/invoices-api/internal/application/interfaces"
	"github.com/pensk/invoices-api/internal/interface/api/handler/request"
)

type UserHandler struct {
	service interfaces.UserService
}

func NewUserHandler(router chi.Router, service interfaces.UserService) {
	handler := &UserHandler{
		service: service,
	}

	router.Post("/users/authenticate", handler.Authenticate)
}

func (h *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var req *request.AuthenticateUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd := req.ToAuthenticateUserCommand()

	res, err := h.service.Authenticate(cmd)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
