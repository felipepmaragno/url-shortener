package app

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/felipepmaragno/url-shortener/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, record domain.URLRecord) error
	FindByCode(ctx context.Context, code string) (domain.URLRecord, error)
}

type CodeGenerator interface {
	Generate(ctx context.Context) (string, error)
}

type Service struct {
	repository Repository
	generator  CodeGenerator
	baseURL    string
	now        func() time.Time
}

type ShortenRequest struct {
	LongURL string
}

type ShortenResult struct {
	Code     string
	ShortURL string
	LongURL  string
}

type ResolveResult struct {
	Code    string
	LongURL string
}

func NewService(repository Repository, generator CodeGenerator, baseURL string) *Service {
	return &Service{
		repository: repository,
		generator:  generator,
		baseURL:    strings.TrimRight(baseURL, "/"),
		now:        time.Now,
	}
}

func (s *Service) Shorten(ctx context.Context, request ShortenRequest) (ShortenResult, error) {
	longURL, err := domain.NormalizeLongURL(request.LongURL)
	if err != nil {
		return ShortenResult{}, err
	}

	code, err := s.generator.Generate(ctx)
	if err != nil {
		return ShortenResult{}, fmt.Errorf("generate code: %w", err)
	}

	record := domain.URLRecord{
		Code:      code,
		LongURL:   longURL,
		CreatedAt: s.now().UTC(),
	}
	if err := s.repository.Create(ctx, record); err != nil {
		return ShortenResult{}, fmt.Errorf("create url record: %w", err)
	}

	return ShortenResult{
		Code:     code,
		ShortURL: s.baseURL + "/" + code,
		LongURL:  longURL,
	}, nil
}

func (s *Service) Resolve(ctx context.Context, code string) (ResolveResult, error) {
	record, err := s.repository.FindByCode(ctx, code)
	if err != nil {
		return ResolveResult{}, err
	}

	return ResolveResult{
		Code:    record.Code,
		LongURL: record.LongURL,
	}, nil
}
