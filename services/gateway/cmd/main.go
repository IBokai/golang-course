package main

import (
	collectorpb "distributed-system/api/gen"
	"distributed-system/services/gateway/config"
	"distributed-system/services/gateway/internal/adapter"
	"distributed-system/services/gateway/internal/handler"
	"distributed-system/services/gateway/internal/usecase"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Load()

	conn, err := grpc.NewClient(cfg.CollectorAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to collector: %v", err)
	}
	defer func() {
		conn.Close()
	}()

	collectorClient := collectorpb.NewRepositoryServiceClient(conn)
	collectorAdapter := adapter.New(collectorClient)
	uc := usecase.New(collectorAdapter)
	handler := handler.New(uc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /repos/{owner}/{name}", handler.GetRepositoryInformation)

	log.Printf("Gateway listening on %s", cfg.HTTPAddr)

	if err := http.ListenAndServe(cfg.HTTPAddr, mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
