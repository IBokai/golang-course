package usecase

import (
	"context"
	"repo-stat/api/internal/domain"
)

type RepositoryUseCase struct {
	processor domain.ProcessorPort
}

func New(processor domain.ProcessorPort) *RepositoryUseCase {
	return &RepositoryUseCase{processor: processor}
}

func (uc *RepositoryUseCase) GetRepositoryInformation(ctx context.Context, owner, name string) (domain.RepositoryInfo, error) {
	if owner == "" || name == "" {
		return domain.RepositoryInfo{}, domain.ErrInvalidArgument
	}
	return uc.processor.GetRepositoryInformation(ctx, owner, name)
}
