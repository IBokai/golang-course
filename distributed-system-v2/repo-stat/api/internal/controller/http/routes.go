package http

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"log/slog"
	"net/http"
	_ "repo-stat/api/docs"
	"repo-stat/api/internal/usecase"
)

func AddRoutes(mux *http.ServeMux, log *slog.Logger, ping *usecase.Ping, repository *usecase.RepositoryUseCase) {
	mux.Handle("GET /api/ping", NewPingHandler(log, ping))
	mux.Handle("GET /api/repositories/info", NewRepositoryHandler(repository))
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)
}
