package config

import (
	"log/slog"

	"github.com/htet-29/home-pos/internal/models"
)

type AppConfig struct {
	Addr         string
	Logger       *slog.Logger
	TemplateDir  string
	DataPath     string
	TemplateData *models.TemplateData
}
