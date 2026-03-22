package handler

import (
	"context"
	collectorpb "distributed-system/api/gen"
	"distributed-system/services/collector/internal/domain"
	"distributed-system/services/collector/internal/usecase"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCHandler struct {
	collectorpb.UnimplementedRepositoryServiceServer
	uc *usecase.RepositoryUseCase
}

func New(uc *usecase.RepositoryUseCase) *GRPCHandler {
	return &GRPCHandler{uc: uc}
}

func (h *GRPCHandler) GetRepositoryInformation(ctx context.Context, req *collectorpb.RepositoryRequest) (*collectorpb.RepositoryResponse, error) {
	repo, err := h.uc.GetRepositoryInformation(ctx, req.Owner, req.RepositoryName)
	if err != nil {
		return nil, mapError(err)
	}
	return &collectorpb.RepositoryResponse{
		Owner:           &collectorpb.RepositoryOwner{Name: repo.Owner.Login},
		Name:            repo.Name,
		Description:     repo.Description,
		StargazersCount: int32(repo.StargazersCount),
		ForksCount:      int32(repo.ForksCount),
		CreationDate:    timestamppb.New(repo.CreationDate),
	}, nil
}

func mapError(err error) error {
	switch {
	case errors.Is(err, domain.ErrRepositoryNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrRateLimitExceeded):
		return status.Error(codes.ResourceExhausted, err.Error())
	case errors.Is(err, domain.ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	default:
		return status.Error(codes.Internal, "internal server error")
	}
}
