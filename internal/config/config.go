package config

import "os"

type Config struct {
	Addr    string
	BaseURL string
}

func Load() Config {
	return Config{
		Addr:    getEnv("ADDR", ":8080"),
		BaseURL: getEnv("BASE_URL", "http://localhost:8080"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
