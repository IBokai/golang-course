package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"repo-stat/collector/internal/domain"
	"time"
)

type githubResponse struct {
	Owner struct {
		Login string `json:"login"`
	}
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	StargazersCount int       `json:"stargazers_count"`
	ForksCount      int       `json:"forks_count"`
	CreationDate    time.Time `json:"created_at"`
}

func (g githubResponse) toDomain() domain.RepositoryInfo {
	return domain.RepositoryInfo{
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
		httpClient: &http.Client{Timeout: 15 * time.Second},
		baseURL:    "https://api.github.com",
	}
}

func (c *GitHubClient) GetRepositoryInformation(ctx context.Context, owner, name string) (domain.RepositoryInfo, error) {
	url := fmt.Sprintf("%s/repos/%s/%s", c.baseURL, owner, name)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return domain.RepositoryInfo{}, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return domain.RepositoryInfo{}, domain.ErrRepositoryNotFound
	case http.StatusForbidden, http.StatusTooManyRequests:
		return domain.RepositoryInfo{}, domain.ErrRateLimitExceeded
	case http.StatusOK:
	default:
		return domain.RepositoryInfo{}, domain.ErrInternalError
	}

	var githubResponse githubResponse
	if err := json.NewDecoder(resp.Body).Decode(&githubResponse); err != nil {
		return domain.RepositoryInfo{}, err
	}
	return githubResponse.toDomain(), nil
}
