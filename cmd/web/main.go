package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))

	mux := http.NewServeMux()

	// mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /home/view", homeView)
	mux.HandleFunc("GET /home/create", homeCreate)
	mux.HandleFunc("POST /home/create", homeCreatePost)

	logger.Info("starting servers", slog.String("addr", *addr))

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
