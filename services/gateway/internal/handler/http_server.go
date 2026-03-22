package handler

import (
	"distributed-system/services/gateway/internal/domain"
	"distributed-system/services/gateway/internal/usecase"
	"encoding/json"
	"errors"
	"net/http"
)

type HTTPHandler struct {
	uc *usecase.RepositoryUseCase
}

func New(uc *usecase.RepositoryUseCase) *HTTPHandler {
	return &HTTPHandler{uc: uc}
}

func (h *HTTPHandler) GetRepositoryInformation(w http.ResponseWriter, r *http.Request) {
	owner := r.PathValue("owner")
	name := r.PathValue("name")
	repo, err := h.uc.GetRepositoryInformation(r.Context(), owner, name)
	if err != nil {
		mapError(w, err)
		return
	}
	json.NewEncoder(w).Encode(repo)
}

func mapError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrRepositoryNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, domain.ErrInvalidArgument):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
