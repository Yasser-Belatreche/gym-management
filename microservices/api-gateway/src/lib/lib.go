package lib

import (
	"gym-management-api-gateway/src/lib/logger"
)

func Logger() logger.Logger {
	return logger.NewLogger()
}
