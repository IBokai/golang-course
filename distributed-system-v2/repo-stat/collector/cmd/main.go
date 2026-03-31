package main

import (
	"repo-stat/collector/internal/app"
	"repo-stat/collector/internal/config"
)

func main() {
	cfg := config.MustLoad()
	app.New(cfg).Run()
}
