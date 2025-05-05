package bootstrap

import (
	"log/slog"
	"os"
	"strconv"
	"url-shortener/internal/config/cleanenv"
	libLoggerSlog "url-shortener/internal/logger/slog"
	"url-shortener/internal/storage/postgresql"
)

type Logger interface {
	Info(msg string, args ...any)
}

type Storage interface {
	StorageWriter
	StorageReader
}

type StorageWriter interface {
	SaveUrl(url, alice string) error
	DeleteUrl(url string) error
}

type StorageReader interface {
	GetUrl(alice string) (string, error)
}

type App struct {
	cfg     *cleanenv.Config
	log     Logger
	storage Storage
}

func Factory() *App {
	cfg := cleanenv.Factory()

	log := libLoggerSlog.MustSetupLogger(cfg.App)
	log = log.With(
		slog.String("program", "main"),
		slog.String("pid", strconv.Itoa(os.Getpid())),
		slog.String("env", cfg.App))
	log.Info("starting server url-shortener")
	log.Debug("debug on")

	storagePostgres, err := postgresql.FactoryStorage(cfg)
	if err != nil {
		log.Error("failed to connect to storage", err.Error())
		os.Exit(1)
	}
	log.Info("connected to db success")

	return &App{cfg: cfg, log: log, storage: storagePostgres}
}

func (app *App) GetStorage() Storage {
	return app.storage
}

func (app *App) GetLogger() Logger {
	return app.log
}
