package main

import (
	"service/internal/app/http_app"
)

func main() {
	// @title Market API
	// @version 1.0
	// @description This is a Market API server.

	// @BasePath /5

	// securityDefinitions.apiKey ApiKeyAuth
	// @in header
	// @name Access
	// @name ua
	app_ := http_app.NewHttpApp(false)
	app_.Run()
}
