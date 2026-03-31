package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"repo-stat/api/internal/domain"
	"repo-stat/api/internal/usecase"
	"strings"
)

func NewRepositoryHandler(uc *usecase.RepositoryUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawURL := r.URL.Query().Get("url")
		if rawURL == "" {
			http.Error(w, "url is required", http.StatusBadRequest)
			return
		}
		owner, name, err := parseGitHubURL(rawURL)
		if err != nil {
			http.Error(w, "invalid github url", http.StatusBadRequest)
			return
		}
		resp, err := uc.GetRepositoryInformation(r.Context(), owner, name)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrRepositoryNotFound):
				http.Error(w, "repository not found", http.StatusNotFound)
			case errors.Is(err, domain.ErrInvalidArgument):
				http.Error(w, "invalid argument", http.StatusBadRequest)
			default:
				http.Error(w, "internal error", http.StatusInternalServerError)
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
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
