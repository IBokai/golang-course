package main

import (
	"repo-stat/processor/internal/app"
	"repo-stat/processor/internal/config"
)

func main() {
	cfg := config.MustLoad() 
	app.New(cfg).Run()
}