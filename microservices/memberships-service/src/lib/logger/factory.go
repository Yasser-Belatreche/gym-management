package logger

import "gym-management-memberships/src/lib/logger/printer"

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
