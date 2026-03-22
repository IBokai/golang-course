package main

import (
	"distributed-system/services/collector/internal/app"
	"distributed-system/services/collector/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.New(cfg).Run()
}
