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
		GRPCAddr: getEnv("GRPC_ADDR", ":50051"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
