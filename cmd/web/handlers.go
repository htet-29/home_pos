package main

import (
	"net/http"
)

func homeView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a home view."))
}

func homeCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a home create."))
}

func homeCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a home data..."))
}
