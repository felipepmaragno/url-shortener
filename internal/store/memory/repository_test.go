package memory

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/felipepmaragno/url-shortener/internal/domain"
)

func TestRepositoryCreateAndFindByCode(t *testing.T) {
	repository := NewRepository()
	record := domain.URLRecord{
		Code:      "abcde",
		LongURL:   "https://example.com",
		CreatedAt: time.Now(),
	}

	if err := repository.Create(context.Background(), record); err != nil {
		t.Fatalf("Create() error = %v, want nil", err)
	}

	got, err := repository.FindByCode(context.Background(), record.Code)
	if err != nil {
		t.Fatalf("FindByCode() error = %v, want nil", err)
	}
	if got != record {
		t.Fatalf("FindByCode() = %+v, want %+v", got, record)
	}
}

func TestRepositoryFindByCodeMissing(t *testing.T) {
	repository := NewRepository()

	_, err := repository.FindByCode(context.Background(), "missing")
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("FindByCode() error = %v, want %v", err, domain.ErrNotFound)
	}
}
