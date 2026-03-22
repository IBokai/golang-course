package config

import "os"

type Config struct {
	HTTPAddr      string
	CollectorAddr string
}

func Load() *Config {
	return &Config{
		HTTPAddr:      getEnv("HTTP_ADDR", ":8080"),
		CollectorAddr: getEnv("COLLECTOR_ADDR", "localhost:50051"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
