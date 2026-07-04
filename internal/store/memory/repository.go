package memory

import (
	"context"
	"sync"

	"github.com/felipepmaragno/url-shortener/internal/domain"
)

type Repository struct {
	mu      sync.RWMutex
	records map[string]domain.URLRecord
}

func NewRepository() *Repository {
	return &Repository{
		records: make(map[string]domain.URLRecord),
	}
}

func (r *Repository) Create(_ context.Context, record domain.URLRecord) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.records[record.Code] = record
	return nil
}

func (r *Repository) FindByCode(_ context.Context, code string) (domain.URLRecord, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	record, ok := r.records[code]
	if !ok {
		return domain.URLRecord{}, domain.ErrNotFound
	}

	return record, nil
}
