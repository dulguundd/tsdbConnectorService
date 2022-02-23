package main

import (
	"github.com/dulguundd/logError-lib/logger"
	"tsdbConnectorService/app"
)

func main() {
	logger.Info("Starting the application.....")
	app.Start()
	logger.Info("Running")
}
