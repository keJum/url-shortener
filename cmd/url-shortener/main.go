package main

import (
	"url-shortener/internal/app/bootstrap"
)

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
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

func main() {
	app := bootstrap.Factory()
	var storage Storage
	var logger Logger
	storage = app.Storage
	logger = app.Log
	_ = storage
	_ = logger

	//logger := app.GetLogger()
	//logger.Info(app.GetStorage().GetUrl("google1"))
}
