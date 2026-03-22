package main

import (
	"distributed-system/services/gateway/internal/app"
	"distributed-system/services/gateway/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.New(cfg).Run()
}
