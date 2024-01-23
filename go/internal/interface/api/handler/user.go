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
	router.Post("/users/signup", handler.Create)
}

func (h *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var req *request.AuthenticateUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd, err := req.ToAuthenticateUserCommand()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.service.Authenticate(cmd)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req *request.CreateUserRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd, err := req.ToCreateUserCommand()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.service.Create(cmd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
