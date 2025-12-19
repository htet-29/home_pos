package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /pos/stock", app.stocksView)
	mux.HandleFunc("GET /pos/stock/create", app.stockCreate)
	mux.HandleFunc("POST /pos/stock/create", app.stockCreatePost)

	return mux
}
