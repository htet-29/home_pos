package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "home.gohtml", templateData{})
}

func (app *application) stocksView(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "stock.gohtml", templateData{})
}

func (app *application) stockCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a stock create."))
}

func (app *application) stockCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a stock data..."))
}
