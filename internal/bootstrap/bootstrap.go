package bootstrap

import (
	"log/slog"
	"os"
	"strconv"
	"url-shortener/internal/bootstrap/internal/config"
	libLoggerSlog "url-shortener/internal/bootstrap/internal/logger/slog"
	"url-shortener/internal/bootstrap/internal/storage/postgresql"
)

func App() *BoostrapStruct {
	cfg := config.MustLoad()

	log := libLoggerSlog.MustSetupLogger(cfg.App)
	log = log.With(
		slog.String("program", "main"),
		slog.String("pid", strconv.Itoa(os.Getpid())),
		slog.String("env", cfg.App))
	log.Info("starting server url-shortener")
	log.Debug("debug on")

	storage, err := postgresql.New(&cfg.Storage)
	if err != nil {
		log.Error("failed to connect to storage", err.Error())
		os.Exit(1)
	}
	log.Info("connected to db success")

	return &BoostrapStruct{Cfg: cfg, Log: log, Storage: storage}
}
