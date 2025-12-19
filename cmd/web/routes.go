package main

import (
	"net/http"

	"github.com/htet-29/home-pos/internal/handlers"
)

func routes(h *handlers.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", h.Home)
	mux.HandleFunc("GET /pos/stock", h.StocksView)
	mux.HandleFunc("GET /pos/stock/create", h.StockCreate)
	mux.HandleFunc("POST /pos/stock/create", h.StockCreatePost)

	return mux
}
