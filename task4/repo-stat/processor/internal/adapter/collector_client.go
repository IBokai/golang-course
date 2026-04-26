package adapter

import (
	"context"
	"repo-stat/processor/internal/domain"
	collectorpb "repo-stat/proto/collector"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CollectorClient struct {
	client collectorpb.CollectorServiceClient
}

func New(conn *grpc.ClientConn) *CollectorClient {
	return &CollectorClient{client: collectorpb.NewCollectorServiceClient(conn)}
}

func (c *CollectorClient) GetRepositoryInformation(ctx context.Context, owner, name string) (domain.RepositoryInfo, error) {
	resp, err := c.client.GetRepositoryInformation(ctx, &collectorpb.RepositoryRequest{Owner: owner, RepositoryName: name})
	if err != nil {
		switch status.Code(err) {
		case codes.NotFound:
			return domain.RepositoryInfo{}, domain.ErrRepositoryNotFound
		case codes.ResourceExhausted:
			return domain.RepositoryInfo{}, domain.ErrRateLimitExceeded
		case codes.InvalidArgument:
			return domain.RepositoryInfo{}, domain.ErrInvalidArgument
		default:
			return domain.RepositoryInfo{}, domain.ErrInternalError
		}
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
