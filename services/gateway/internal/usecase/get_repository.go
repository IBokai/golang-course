package usecase

import (
	"context"
	"distributed-system/services/gateway/internal/domain"
)

type RepositoryUseCase struct {
	collector domain.CollectorPort
}

func New(collector domain.CollectorPort) *RepositoryUseCase {
	return &RepositoryUseCase{collector: collector}
}

func (uc *RepositoryUseCase) GetRepositoryInformation(ctx context.Context, owner, name string) (*domain.RepositoryInfo, error) {
	if owner == "" || name == "" {
		return nil, domain.ErrInvalidArgument
	}
	return uc.collector.GetRepositoryInformation(ctx, owner, name)
}
