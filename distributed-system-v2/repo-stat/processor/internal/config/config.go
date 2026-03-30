package config

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	GRPCAddr      string
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
	grpcAddr := os.Getenv("PROCESSOR_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, errors.New("PROCESSOR_GRPC_ADDR is required")
	}
	collectorAddr := os.Getenv("COLLECTOR_ADDR")
	if collectorAddr == "" {
		return nil, errors.New("COLLECTOR_ADDR is required")
	}
	return &Config{
		GRPCAddr:      grpcAddr,
		CollectorAddr: collectorAddr,
	}, nil
}
