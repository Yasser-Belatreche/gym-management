package logger

import (
	"encoding/json"
	"gym-management/src/lib/logger/printer"
	"gym-management/src/lib/primitives/application_specific"
	"time"
)

type Facade struct {
	printer printer.Printer
}

func (f *Facade) Info(msg string, payload map[string]interface{}, session *application_specific.Session) {
	date := time.Now().Format(time.RFC3339)

	payloadStr := ""

	if payload != nil {
		bytes, err := json.Marshal(payload)
		if err == nil {
			payloadStr = string(bytes)
		}
	}

	str := session.CorrelationId + " - " + date + " - " + "INFO" + " - " + msg + " - " + "Payload: " + payloadStr

	f.printer.Print(str)
}

func (f *Facade) Warn(msg string, payload map[string]interface{}, error *error, session *application_specific.Session) {
	date := time.Now().Format(time.RFC3339)

	payloadStr := ""

	if payload != nil {
		bytes, err := json.Marshal(payload)
		if err == nil {
			payloadStr = string(bytes)
		}
	}

	errorStr := ""
	if error != nil {
		errorStr = (*error).Error()
	}

	str := session.CorrelationId + " - " + date + " - " + "WARN" + " - " + msg + " - " + "Payload: " + payloadStr + " - " + "Error: " + errorStr

	f.printer.Print(str)
}

func (f *Facade) Error(msg string, payload map[string]interface{}, error *error, session *application_specific.Session) {
	date := time.Now().Format(time.RFC3339)

	payloadStr := ""

	if payload != nil {
		bytes, err := json.Marshal(payload)
		if err == nil {
			payloadStr = string(bytes)
		}
	}

	errorStr := ""

	if error != nil {
		errorStr = (*error).Error()
	}

	str := session.CorrelationId + " - " + date + " - " + "ERROR" + " - " + msg + " - " + "Payload: " + payloadStr + " - " + "Error: " + errorStr

	f.printer.Print(str)
}

func (f *Facade) Clear() {
	f.printer.Clear()
}
