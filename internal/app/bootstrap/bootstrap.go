package bootstrap

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"os"
	"strconv"
	"url-shortener/internal/config/cleanenv"
	middlewareLogger "url-shortener/internal/http-server/middleware/logger"
	libLoggerSlog "url-shortener/internal/logger/slog"
	"url-shortener/internal/storage/postgresql"
)

type App struct {
	Cfg     *cleanenv.Config
	Log     *slog.Logger
	Storage *postgresql.Storage
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

	storage, err := postgresql.FactoryStorage(cfg)
	if err != nil {
		log.Error("failed to connect to storage", err.Error())
		os.Exit(1)
	}
	log.Info("connected to db success")

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middlewareLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	return &App{Cfg: cfg, Log: log, Storage: storage}
}
