package main

import (
	"net/http"

	"github.com/htet-29/home-pos/ui"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /pos/stock", app.stocksView)
	mux.HandleFunc("GET /pos/stock/create", app.stockCreate)
	mux.HandleFunc("POST /pos/stock/create", app.stockCreatePost)

	return mux
}
