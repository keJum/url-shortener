package main

import (
	"url-shortener/internal/app/bootstrap"
	"url-shortener/internal/lib/logger/slog"
)

func main() {
	app := bootstrap.Factory()
	_ = app
	storage := app.Storage
	log := app.Log
	if err := storage.SaveUrl("https://google.com", "google.com"); err != nil {
		log.Error("failed to save url", slog.Err(err))
	}
	if err := storage.SaveUrl("https://google.com", "goog.com"); err != nil {
		log.Error("failed to save url", slog.Err(err))
	}
}
