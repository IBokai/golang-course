package usecase

import (
	"context"
	"repo-stat/processor/internal/domain"
)

type RepositoryUseCase struct {
	collector domain.CollectorPort
}

func NewRepositoryUsecase(collector domain.CollectorPort) *RepositoryUseCase {
	return &RepositoryUseCase{collector: collector}
}

func (uc *RepositoryUseCase) GetRepositoryInformation(ctx context.Context, owner, name string) (domain.RepositoryInfo, error) {
	if owner == "" || name == "" {
		return domain.RepositoryInfo{}, domain.ErrInvalidArgument
	}
	return uc.collector.GetRepositoryInformation(ctx, owner, name)
}