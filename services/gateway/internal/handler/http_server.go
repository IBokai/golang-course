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

// GetRepositoryInforamtion godoc
// @Summary Get information about a repository
// @Description Returns information about a GitHub repository
// @Tags repository
// @Produce json
// @Param owner path string true "Repository owner"
// @Param name path string true "Repository name"
// @Success 200 {object} domain.RepositoryInfo
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /repos/{owner}/{name} [get]
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
