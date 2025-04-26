package bootstrap

import (
	"log/slog"
	"url-shortener/internal/bootstrap/internal/config"
	"url-shortener/internal/bootstrap/internal/storage/postgresql"
)

type BoostrapStruct struct {
	Cfg     *config.Config
	Log     *slog.Logger
	Storage *postgresql.Storage
}

type BoostrapInterface interface {
	App() *BoostrapStruct
}
