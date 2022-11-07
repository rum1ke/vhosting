package messages

import (
	"vhosting/pkg/logger"
)

func FatalFailedToLoadEnvironmentFile(err error) *logger.Log {
	return &logger.Log{ErrCode: 5, Message: "Failed to load environment file. Error: " + err.Error(), ErrLevel: logger.ErrLevelFatal}
}

func InfoEnvironmentsLoaded() *logger.Log {
	return &logger.Log{Message: "Environments loaded"}
}
