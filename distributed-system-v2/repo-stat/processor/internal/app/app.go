package app

import (
	"log"
	"net"
	"repo-stat/processor/internal/adapter"
	"repo-stat/processor/internal/config"
	"repo-stat/processor/internal/handler"
	"repo-stat/processor/internal/usecase"
	processorbp "repo-stat/proto/processor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	server *grpc.Server
	lis    net.Listener
}

func New(cfg *config.Config) *App {
	collectorConn, err := grpc.NewClient(cfg.CollectorAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to collector: %v", err)
	}
	collectorClient := adapter.New(collectorConn)
	uc := usecase.NewRepositoryUsecase(collectorClient)
	handler := handler.New(uc)
	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to lister: %v", err)
	}
	server := grpc.NewServer()
	processorbp.RegisterProcessorServiceServer(server, handler)
	return &App{
		server: server,
		lis:    lis,
	}
}

func (a *App) Run() {
	log.Printf("Processor listening on %s", a.lis.Addr())
	if err := a.server.Serve(a.lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
