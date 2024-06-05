package logger

import "gym-management-api-gateway/src/lib/logger/printer"

var instance Logger

func NewLogger() Logger {
	if instance != nil {
		return instance
	}

	instance = &Facade{
		printer: &printer.StdoutPrinter{},
	}

	return instance
}
