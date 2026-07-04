package app

import (
	"context"
	"errors"
	"testing"

	"github.com/felipepmaragno/url-shortener/internal/codegen"
	"github.com/felipepmaragno/url-shortener/internal/domain"
	"github.com/felipepmaragno/url-shortener/internal/store/memory"
)

func TestServiceShortenAndResolve(t *testing.T) {
	service := NewService(memory.NewRepository(), codegen.NewSequentialGenerator(), "http://localhost:8080")

	created, err := service.Shorten(context.Background(), ShortenRequest{LongURL: "https://example.com/path"})
	if err != nil {
		t.Fatalf("Shorten() error = %v, want nil", err)
	}

	if created.Code == "" {
		t.Fatal("Shorten() code is empty")
	}
	if created.LongURL != "https://example.com/path" {
		t.Fatalf("Shorten() long URL = %q, want %q", created.LongURL, "https://example.com/path")
	}
	if created.ShortURL != "http://localhost:8080/"+created.Code {
		t.Fatalf("Shorten() short URL = %q, want %q", created.ShortURL, "http://localhost:8080/"+created.Code)
	}

	resolved, err := service.Resolve(context.Background(), created.Code)
	if err != nil {
		t.Fatalf("Resolve() error = %v, want nil", err)
	}
	if resolved.LongURL != created.LongURL {
		t.Fatalf("Resolve() long URL = %q, want %q", resolved.LongURL, created.LongURL)
	}
}

func TestServiceShortenRejectsInvalidURL(t *testing.T) {
	service := NewService(memory.NewRepository(), codegen.NewSequentialGenerator(), "http://localhost:8080")

	_, err := service.Shorten(context.Background(), ShortenRequest{LongURL: "example.com"})
	if !errors.Is(err, domain.ErrInvalidURL) {
		t.Fatalf("Shorten() error = %v, want %v", err, domain.ErrInvalidURL)
	}
}

func TestServiceResolveMissingCode(t *testing.T) {
	service := NewService(memory.NewRepository(), codegen.NewSequentialGenerator(), "http://localhost:8080")

	_, err := service.Resolve(context.Background(), "missing")
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("Resolve() error = %v, want %v", err, domain.ErrNotFound)
	}
}
