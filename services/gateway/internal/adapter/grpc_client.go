package adapter

import (
	"context"
	 collectorpb "distributed-system/api/gen"
	"distributed-system/services/gateway/internal/domain"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CollectorClient struct {
	client collectorpb.RepositoryServiceClient
}

func New(client collectorpb.RepositoryServiceClient) *CollectorClient {
	return &CollectorClient{client: client}
}

func (c *CollectorClient) GetRepositoryInformation(ctx context.Context, owner, name string) (*domain.RepositoryInfo, error) {
	resp, err := c.client.GetRepositoryInformation(ctx, &collectorpb.RepositoryRequest{Owner: owner, RepositoryName: name})
	if err != nil {
		return nil, mapError(err)
	}
	return &domain.RepositoryInfo{
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
