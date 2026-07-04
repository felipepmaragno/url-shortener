package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/felipepmaragno/url-shortener/internal/app"
	"github.com/felipepmaragno/url-shortener/internal/domain"
)

type Handler struct {
	service *app.Service
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	Code     string `json:"code"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

type errorResponse struct {
	Error apiError `json:"error"`
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewHandler(service *app.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", h.health)
	mux.HandleFunc("POST /api/v1/shorten", h.shorten)
	mux.HandleFunc("GET /", h.redirect)
	return mux
}

func (h *Handler) health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok\n"))
}

func (h *Handler) shorten(w http.ResponseWriter, r *http.Request) {
	var request shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "invalid JSON request body")
		return
	}

	result, err := h.service.Shorten(r.Context(), app.ShortenRequest{LongURL: request.URL})
	if err != nil {
		h.writeMappedError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, shortenResponse{
		Code:     result.Code,
		ShortURL: result.ShortURL,
		LongURL:  result.LongURL,
	})
}

func (h *Handler) redirect(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		http.NotFound(w, r)
		return
	}

	result, err := h.service.Resolve(r.Context(), code)
	if err != nil {
		h.writeMappedError(w, err)
		return
	}

	http.Redirect(w, r, result.LongURL, http.StatusFound)
}

func (h *Handler) writeMappedError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidURL):
		writeError(w, http.StatusBadRequest, "invalid_url", "invalid URL")
	case errors.Is(err, domain.ErrNotFound):
		writeError(w, http.StatusNotFound, "not_found", "short URL not found")
	default:
		writeError(w, http.StatusInternalServerError, "internal_error", "internal server error")
	}
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, errorResponse{
		Error: apiError{
			Code:    code,
			Message: message,
		},
	})
}
