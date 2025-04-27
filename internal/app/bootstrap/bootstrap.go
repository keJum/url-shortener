package bootstrap

import (
	"log/slog"
	"os"
	"strconv"
	"url-shortener/internal/config"
	"url-shortener/internal/config/cleanenv"
	libLoggerSlog "url-shortener/internal/logger/slog"
	"url-shortener/internal/storage"
	"url-shortener/internal/storage/postgresql"
)

type App struct {
	Cfg     config.Config
	Log     *slog.Logger
	Storage storage.Storage
}

func Factory() *App {
	cfg := cleanenv.Factory()

	log := libLoggerSlog.MustSetupLogger(cfg.GetApp())
	log = log.With(
		slog.String("program", "main"),
		slog.String("pid", strconv.Itoa(os.Getpid())),
		slog.String("env", cfg.GetApp()))
	log.Info("starting server url-shortener")
	log.Debug("debug on")

	storagePostgres, err := postgresql.FactoryStorage(cfg.GetStorage())
	if err != nil {
		log.Error("failed to connect to storage", err.Error())
		os.Exit(1)
	}
	log.Info("connected to db success")

	return &App{Cfg: cfg, Log: log, Storage: storagePostgres}
}
