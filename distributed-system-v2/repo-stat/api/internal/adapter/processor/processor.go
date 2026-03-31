package processor

import (
	"context"
	"errors"
	"repo-stat/api/internal/domain"
	processorpb "repo-stat/proto/processor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProcessorClient struct {
	client processorpb.ProcessorServiceClient
}

func New(conn *grpc.ClientConn) *ProcessorClient {
	return &ProcessorClient{client: processorpb.NewProcessorServiceClient(conn)}
}

func (c *ProcessorClient) Ping(ctx context.Context) domain.PingStatus {
	_, err := c.client.Ping(ctx, &processorpb.PingRequest{})
	if err != nil {
		return domain.PingStatusDown
	}
	return domain.PingStatusUp
}

func (c *ProcessorClient) GetRepositoryInformation(ctx context.Context, owner, name string) (domain.RepositoryInfo, error) {
	resp, err := c.client.GetRepositoryInformation(ctx, &processorpb.RepositoryRequest{Owner: owner, RepositoryName: name})
	if err != nil {
		return domain.RepositoryInfo{}, mapError(err)
	}
	return domain.RepositoryInfo{
		Owner:           domain.Owner{Login: resp.Owner.Name},
		Name:            resp.Name,
		Description:     resp.Description,
		StargazersCount: int(resp.StargazersCount),
		ForksCount:      int(resp.ForksCount),
		CreationDate:    resp.CreationDate.AsTime(),
	}, nil
}

func mapError(err error) error {
	st, ok := status.FromError(err)
	if !ok {
		return err
	}
	switch st.Code() {
	case codes.NotFound:
		return domain.ErrRepositoryNotFound
	case codes.InvalidArgument:
		return domain.ErrInvalidArgument
	default:
		return errors.New("internal error")
	}
}
