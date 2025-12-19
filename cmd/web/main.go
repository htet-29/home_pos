package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/htet-29/home-pos/internal/config"
	"github.com/htet-29/home-pos/internal/handlers"
)

func main() {
	addr := flag.String("addr", "localhost:4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	app := &config.AppConfig{
		Addr:   *addr,
		Logger: logger,
	}

	app.Logger.Info("starting servers", slog.String("addr", *addr))

	h := handlers.NewHandler(app)

	err := http.ListenAndServe(*addr, routes(h))
	app.Logger.Error(err.Error())
	os.Exit(1)
}
