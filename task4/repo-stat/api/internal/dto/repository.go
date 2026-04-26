package dto

import (
	"repo-stat/api/internal/domain"
	"time"
)

type RepositoryResponse struct {
	FullName    string `json:"full_name"`
	Description string `json:"description"`
	Stars       int64  `json:"stars"`
	Forks       int64  `json:"forks"`
	CreatedAt   string `json:"created_at"`
}

func NewRepositoryResponse(info domain.RepositoryInfo) RepositoryResponse {
	return RepositoryResponse{
		FullName:    info.Owner.Login + "/" + info.Name,
		Description: info.Description,
		Stars:       int64(info.StargazersCount),
		Forks:       int64(info.ForksCount),
		CreatedAt:   info.CreationDate.UTC().Format(time.RFC3339),
	}
}
