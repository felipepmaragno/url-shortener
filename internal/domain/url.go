package domain

import (
	"errors"
	"net/url"
	"strings"
	"time"
)

var (
	ErrInvalidURL = errors.New("invalid url")
	ErrNotFound   = errors.New("url not found")
)

type URLRecord struct {
	Code      string
	LongURL   string
	CreatedAt time.Time
}

func NormalizeLongURL(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", ErrInvalidURL
	}

	parsed, err := url.ParseRequestURI(trimmed)
	if err != nil {
		return "", ErrInvalidURL
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", ErrInvalidURL
	}
	if parsed.Host == "" {
		return "", ErrInvalidURL
	}

	return parsed.String(), nil
}
