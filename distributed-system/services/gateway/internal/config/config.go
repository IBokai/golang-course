package config

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	HTTPAddr      string
	CollectorAddr string
}

func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}

func Load() (*Config, error) {
	httpAddr := os.Getenv("HTTP_ADDR")
	if httpAddr == "" {
		return nil, errors.New("HTTP_ADDR is required")
	}
	collectorAddr := os.Getenv("COLLECTOR_ADDR")
	if collectorAddr == "" {
		return nil, errors.New("COLLECTOR_ADDR is required")
	}
	return &Config{
		HTTPAddr:      httpAddr,
		CollectorAddr: collectorAddr,
	}, nil
}
