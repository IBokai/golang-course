package app

import (
	collectorpb "distributed-system/api/gen"
	_ "distributed-system/services/gateway/docs"
	"distributed-system/services/gateway/internal/adapter"
	"distributed-system/services/gateway/internal/config"
	"distributed-system/services/gateway/internal/handler"
	"distributed-system/services/gateway/internal/usecase"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	server *http.Server
	conn   *grpc.ClientConn
}

func New(cfg *config.Config) *App {
	conn, err := grpc.NewClient(cfg.CollectorAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to collector: %v", err)
	}

	collectorClient := collectorpb.NewRepositoryServiceClient(conn)
	collectorAdapter := adapter.New(collectorClient)
	uc := usecase.New(collectorAdapter)
	handler := handler.New(uc)

	mux := http.NewServeMux()
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("GET /repos/{owner}/{name}", handler.GetRepositoryInformation)

	return &App{
		server: &http.Server{
			Addr:    cfg.HTTPAddr,
			Handler: mux,
		},
	}
}

func (a *App) Run() {
	defer func() {
		_ = a.conn.Close()
	}()

	log.Printf("Gateway listening on %s", a.server.Addr)

	if err := a.server.ListenAndServe(); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
