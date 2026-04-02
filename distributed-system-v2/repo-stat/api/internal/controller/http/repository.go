package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"repo-stat/api/internal/domain"
	"repo-stat/api/internal/dto"
	"repo-stat/api/internal/usecase"
	"strings"
)

// @Summary Get repository info
// @Description Get GitHub repository information by URL
// @Tags repositories
// @Param url query string true "GitHub repository URL"
// @Produce json
// @Success 200 {object} domain.RepositoryInfo
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /api/repositories/info [get]
func NewRepositoryHandler(uc *usecase.RepositoryUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawURL := r.URL.Query().Get("url")
		if rawURL == "" {
			writeError(w, "url is required", http.StatusBadRequest)
			return
		}
		owner, name, err := parseGitHubURL(rawURL)
		if err != nil {
			writeError(w, "invalid github url", http.StatusBadRequest)
			return
		}
		resp, err := uc.GetRepositoryInformation(r.Context(), owner, name)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrRepositoryNotFound):
				writeError(w, "repository not found", http.StatusNotFound)
			case errors.Is(err, domain.ErrInvalidArgument):
				writeError(w, "invalid argument", http.StatusBadRequest)
			default:
				writeError(w, "internal error", http.StatusInternalServerError)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dto.NewRepositoryResponse(resp)); err != nil {
			log.Printf("failed to write response")
		}
	}
}

func parseGitHubURL(rawURL string) (owner, name string, err error) {
	rawURL = strings.TrimRight(rawURL, "/")
	parts := strings.Split(rawURL, "/")
	if len(parts) < 2 {
		return "", "", domain.ErrInvalidArgument
	}
	return parts[len(parts)-2], parts[len(parts)-1], nil
}

func writeError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": msg}); err != nil {
		log.Printf("failed to write error response")
	}
}
