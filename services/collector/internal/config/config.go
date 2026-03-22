package config

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	GRPCAddr string
}

func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}

func Load() (*Config, error) {
	grpcAddr := os.Getenv("GRPC_ADDR")
	if grpcAddr == "" {
		return nil, errors.New("GRPC_ADDR is required")
	}
	return &Config{
		GRPCAddr: grpcAddr,
	}, nil
}

