package domain

import (
	"context"
	"time"
)

type Owner struct {
	Login string
}

type RepositoryInfo struct {
	Owner           Owner
	Name            string
	Description     string
	StargazersCount int
	ForksCount      int
	CreationDate    time.Time
}

type GitHubPort interface {
	GetRepositoryInformation(ctx context.Context, owner, name string) (*RepositoryInfo, error)
}
