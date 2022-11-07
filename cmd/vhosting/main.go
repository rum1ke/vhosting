package main

import (
	"vhosting/internal/messages"
	"vhosting/pkg/config"
	logger "vhosting/pkg/logger"
	"vhosting/pkg/server"

	"github.com/joho/godotenv"
)

const configsDir string = "./configs/"

func main() {
	// Read .env file.
	err := godotenv.Load(configsDir + ".env")
	if err != nil {
		logger.Print(messages.FatalFailedToLoadEnvironmentFile(err))
		return
	}
	logger.Print(messages.InfoEnvironmentsLoaded())

	// Load config file.
	cfg, err := config.LoadConfig(configsDir + "config.yml")
	if err != nil {
		logger.Print(messages.FatalFailedToLoadConfigFile(err))
		return
	}
	logger.Print(messages.InfoConfigLoaded())

	// Get instance of HTTP server.
	app := server.NewApp(cfg)

	// Run HTTP server.
	err = app.Run()
	if err != nil {
		logger.Print(messages.FatalFailureOnServerRunning(err))
		return
	}
	logger.Print(messages.InfoServerStartedSuccessfullyAtLocalAddress(cfg.ServerHost, cfg.ServerPort))
}
