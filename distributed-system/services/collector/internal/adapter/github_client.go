package adapter

import (
	"context"
	"distributed-system/services/collector/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type githubRepository struct {
	Owner struct {
		Login string `json:"login"`
	}
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	StargazersCount int       `json:"stargazers_count"`
	ForksCount      int       `json:"forks_count"`
	CreationDate    time.Time `json:"created_at"`
}

func (g *githubRepository) toDomain() *domain.RepositoryInfo {
	return &domain.RepositoryInfo{
		Owner:           domain.Owner{Login: g.Owner.Login},
		Name:            g.Name,
		Description:     g.Description,
		StargazersCount: g.StargazersCount,
		ForksCount:      g.ForksCount,
		CreationDate:    g.CreationDate,
	}
}

type GitHubClient struct {
	httpClient *http.Client
	baseURL    string
}

func New() *GitHubClient {
	return &GitHubClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		baseURL:    "https://api.github.com",
	}
}

func (c *GitHubClient) GetRepositoryInformation(ctx context.Context, owner, name string) (*domain.RepositoryInfo, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", c.baseURL, owner, name)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, domain.ErrRepositoryNotFound
	case http.StatusForbidden, http.StatusTooManyRequests:
		return nil, domain.ErrRateLimitExceeded
	}

	var githubResponse githubRepository
	if err := json.NewDecoder(resp.Body).Decode(&githubResponse); err != nil {
		return nil, err
	}
	return githubResponse.toDomain(), nil
}
