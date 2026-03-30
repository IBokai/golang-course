package handler

import (
	"context"
	processorpb "repo-stat/proto/processor"
)

func (h *GRPCHandler) Ping(ctx context.Context, req *processorpb.PingRequest) (*processorpb.PingResponse, error) {
	return &processorpb.PingResponse{Reply: "ok"}, nil
}