package app

import (
	"log"
	"net"
	"repo-stat/collector/internal/adapter"
	"repo-stat/collector/internal/config"
	"repo-stat/collector/internal/handler"
	"repo-stat/collector/internal/usecase"
	collectorpb "repo-stat/proto/collector"

	"google.golang.org/grpc"
)

type App struct {
	server *grpc.Server
	lis net.Listener
}

func New(cfg *config.Config) *App {
	githubClient := adapter.New()
	uc := usecase.New(githubClient)
	grpcHandler := handler.New(uc)

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	collectorpb.RegisterCollectorServiceServer(server, grpcHandler)

	return &App{
		server: server,
		lis:    lis,
	}
}

func (a *App) Run() {
	log.Printf("Collector listening on %s", a.lis.Addr())

	if err := a.server.Serve(a.lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}