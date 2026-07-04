package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/felipepmaragno/url-shortener/internal/app"
	"github.com/felipepmaragno/url-shortener/internal/codegen"
	"github.com/felipepmaragno/url-shortener/internal/store/memory"
)

func TestHandlerHealth(t *testing.T) {
	handler := newTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if strings.TrimSpace(rec.Body.String()) != "ok" {
		t.Fatalf("body = %q, want ok", rec.Body.String())
	}
}

func TestHandlerShortenAndRedirect(t *testing.T) {
	handler := newTestHandler()

	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/shorten", strings.NewReader(`{"url":"https://example.com/path"}`))
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()

	handler.ServeHTTP(createRec, createReq)

	if createRec.Code != http.StatusCreated {
		t.Fatalf("create status = %d, want %d; body=%s", createRec.Code, http.StatusCreated, createRec.Body.String())
	}

	var createResp shortenResponse
	if err := json.NewDecoder(createRec.Body).Decode(&createResp); err != nil {
		t.Fatalf("decode create response: %v", err)
	}
	if createResp.Code == "" {
		t.Fatal("create response code is empty")
	}
	if createResp.ShortURL != "http://localhost:8080/"+createResp.Code {
		t.Fatalf("short_url = %q, want %q", createResp.ShortURL, "http://localhost:8080/"+createResp.Code)
	}

	redirectReq := httptest.NewRequest(http.MethodGet, "/"+createResp.Code, nil)
	redirectRec := httptest.NewRecorder()

	handler.ServeHTTP(redirectRec, redirectReq)

	if redirectRec.Code != http.StatusFound {
		t.Fatalf("redirect status = %d, want %d", redirectRec.Code, http.StatusFound)
	}
	if got := redirectRec.Header().Get("Location"); got != "https://example.com/path" {
		t.Fatalf("Location = %q, want https://example.com/path", got)
	}
}

func TestHandlerShortenInvalidURL(t *testing.T) {
	handler := newTestHandler()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/shorten", strings.NewReader(`{"url":"example.com"}`))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}

	var got errorResponse
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decode error response: %v", err)
	}
	if got.Error.Code != "invalid_url" {
		t.Fatalf("error code = %q, want invalid_url", got.Error.Code)
	}
}

func TestHandlerRedirectMissingCode(t *testing.T) {
	handler := newTestHandler()
	req := httptest.NewRequest(http.MethodGet, "/missing", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

func TestHandlerMethodNotAllowed(t *testing.T) {
	handler := newTestHandler()
	req := httptest.NewRequest(http.MethodPut, "/api/v1/shorten", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
}

func newTestHandler() http.Handler {
	service := app.NewService(memory.NewRepository(), codegen.NewSequentialGenerator(), "http://localhost:8080")
	return NewHandler(service).Routes()
}
