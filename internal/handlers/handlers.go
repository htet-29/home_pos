package handlers

import (
	"net/http"

	"github.com/htet-29/home-pos/internal/config"
	"github.com/htet-29/home-pos/internal/render"
)

type Handler struct {
	app *config.AppConfig
}

func NewHandler(a *config.AppConfig) *Handler {
	return &Handler{
		app: a,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, http.StatusOK, "home.gohtml", h.app)
}

func (h *Handler) StocksView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a stock view."))
}

func (h *Handler) StockCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a stock create."))
}

func (h *Handler) StockCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a stock data..."))
}
