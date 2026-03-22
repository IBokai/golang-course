package app

import (
	collectorpb "distributed-system/api/gen"
	"distributed-system/services/collector/internal/adapter"
	"distributed-system/services/collector/internal/config"
	"distributed-system/services/collector/internal/handler"
	"distributed-system/services/collector/internal/usecase"
	"log"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	server *grpc.Server
	lis    net.Listener
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
	collectorpb.RegisterRepositoryServiceServer(server, grpcHandler)

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
