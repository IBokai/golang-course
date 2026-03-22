package usecase

import (
	"context"
	"distributed-system/services/collector/internal/domain"
)

type RepositoryUseCase struct {
	github domain.GitHubPort
}

func New(github domain.GitHubPort) *RepositoryUseCase {
	return &RepositoryUseCase{github: github}
}

func (uc *RepositoryUseCase) GetRepositoryInformation(ctx context.Context, owner, name string) (*domain.RepositoryInfo, error) {
	if owner == "" || name == "" {
		return nil, domain.ErrInvalidArgument
	}
	return uc.github.GetRepositoryInformation(ctx, owner, name)
}
