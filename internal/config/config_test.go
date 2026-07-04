package config

import "testing"

func TestLoad(t *testing.T) {
	t.Setenv("ADDR", ":9090")
	t.Setenv("BASE_URL", "http://short.local")

	cfg := Load()

	if cfg.Addr != ":9090" {
		t.Fatalf("Addr = %q, want :9090", cfg.Addr)
	}
	if cfg.BaseURL != "http://short.local" {
		t.Fatalf("BaseURL = %q, want http://short.local", cfg.BaseURL)
	}
}

func TestLoadDefaults(t *testing.T) {
	t.Setenv("ADDR", "")
	t.Setenv("BASE_URL", "")

	cfg := Load()

	if cfg.Addr != ":8080" {
		t.Fatalf("Addr = %q, want :8080", cfg.Addr)
	}
	if cfg.BaseURL != "http://localhost:8080" {
		t.Fatalf("BaseURL = %q, want http://localhost:8080", cfg.BaseURL)
	}
}
