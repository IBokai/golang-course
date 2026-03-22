package main

import (
	collectorpb "distributed-system/api/gen"
	"distributed-system/services/collector/config"
	"distributed-system/services/collector/internal/adapter"
	"distributed-system/services/collector/internal/handler"
	"distributed-system/services/collector/internal/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.MustLoad()

	githubClient := adapter.New()
	uc := usecase.New(githubClient)
	grpcHandler := handler.New(uc)
	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	collectorpb.RegisterRepositoryServiceServer(server, grpcHandler)

	log.Printf("Collector listening on %s", cfg.GRPCAddr)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
