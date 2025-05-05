package main

import (
	"url-shortener/internal/app/bootstrap"
)

func main() {
	app := bootstrap.Factory()
	_ = app
	//logger := app.GetLogger()
	//logger.Info(app.GetStorage().GetUrl("google1"))
}
