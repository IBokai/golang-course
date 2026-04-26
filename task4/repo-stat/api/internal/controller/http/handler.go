package http

import (
	"context"
	"log/slog"
	"net/http"
	"repo-stat/api/config"
	"repo-stat/api/internal/adapter/processor"
	"repo-stat/api/internal/adapter/subscriber"
	"repo-stat/api/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHandler(ctx context.Context, log *slog.Logger, cfg config.Config) (http.Handler, error) {
	subscriberClient, err := subscriber.NewClient(cfg.Services.Subscriber, log)
	if err != nil {
		log.Error("cannot init subscriber adapter", "error", err)
		return nil, err
	}

	processorConn, err := grpc.NewClient(cfg.Services.Processor, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("failed to connect to processor", "error", err)
		return nil, err
	}
	processorClient := processor.New(processorConn)
	pingUseCase := usecase.NewPing(subscriberClient, processorClient)
	repositoryUseCase := usecase.New(processorClient)

	mux := http.NewServeMux()
	AddRoutes(mux, log, pingUseCase, repositoryUseCase)

	var handler http.Handler = mux
	return handler, nil
}
