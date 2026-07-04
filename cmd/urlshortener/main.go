package main

import (
	"log"
	"net/http"

	"github.com/felipepmaragno/url-shortener/internal/app"
	"github.com/felipepmaragno/url-shortener/internal/codegen"
	"github.com/felipepmaragno/url-shortener/internal/config"
	"github.com/felipepmaragno/url-shortener/internal/httpapi"
	"github.com/felipepmaragno/url-shortener/internal/store/memory"
)

func main() {
	cfg := config.Load()

	service := app.NewService(memory.NewRepository(), codegen.NewSequentialGenerator(), cfg.BaseURL)
	handler := httpapi.NewHandler(service).Routes()

	log.Printf("starting url-shortener on %s", cfg.Addr)
	if err := http.ListenAndServe(cfg.Addr, handler); err != nil {
		log.Fatal(err)
	}
}
