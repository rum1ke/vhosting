package messages

import (
	"vhosting/pkg/logger"
)

func ErrorCannotDoLogging(err error) *logger.Log {
	return &logger.Log{StatusCode: 500, ErrCode: 400, Message: "Cannot do logging. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}
