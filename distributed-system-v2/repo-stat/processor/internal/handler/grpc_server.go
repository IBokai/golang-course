package handler

import (
	"context"
	"errors"
	"repo-stat/processor/internal/domain"
	"repo-stat/processor/internal/usecase"
	processorpb "repo-stat/proto/processor"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCHandler struct {
	processorpb.UnimplementedProcessorServiceServer
	uc *usecase.RepositoryUseCase
}

func New(uc *usecase.RepositoryUseCase) *GRPCHandler {
	return &GRPCHandler{uc: uc}

}

func (h *GRPCHandler) GetRepositoryInformation(ctx context.Context, req *processorpb.RepositoryRequest) (*processorpb.RepositoryResponse, error) {
	repo, err := h.uc.GetRepositoryInformation(ctx, req.Owner, req.RepositoryName)
	if err != nil {
		return nil, mapError(err)
	}
	return &processorpb.RepositoryResponse{
		Owner:           &processorpb.RepositoryOwner{Name: repo.Owner.Login},
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
