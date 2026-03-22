// @title Repository Info API
// @version 1.0
// @description Getting information about a GitHub repository
// @host localhost:8080
// @BasePath /

package main

import (
	"distributed-system/services/gateway/internal/app"
	"distributed-system/services/gateway/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.New(cfg).Run()
}
