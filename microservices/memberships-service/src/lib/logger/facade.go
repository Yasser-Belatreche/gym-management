package logger

import (
	"encoding/json"
	"gym-management-memberships/src/lib/logger/printer"
	"gym-management-memberships/src/lib/primitives/application_specific"
	"strconv"
	"time"
)

type Facade struct {
	printer printer.Printer
}

func (f *Facade) Info(msg string, payload map[string]interface{}, session *application_specific.Session) {
	now := time.Now()
	epoch := now.UnixNano()
	date := now.Format("2006-01-02 15:04:05.000")

	payloadStr := ""

	if payload != nil {
		bytes, err := json.Marshal(payload)
		if err == nil {
			payloadStr = string(bytes)
		}
	}

	str := session.CorrelationId + " | " + "Date: " + strconv.FormatInt(epoch, 10) + " - " + date + " | " + "INFO" + " | " + msg + " | " + "Payload: " + payloadStr

	f.printer.Print(str)
}

func (f *Facade) Warn(msg string, payload map[string]interface{}, error *error, session *application_specific.Session) {
	now := time.Now()
	epoch := now.UnixNano()
	date := now.Format("2006-01-02 15:04:05.000")

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

	str := session.CorrelationId + " | " + "Date: " + strconv.FormatInt(epoch, 10) + " - " + date + " | " + "WARN" + " | " + msg + " | " + "Payload: " + payloadStr + " | " + "Error: " + errorStr

	f.printer.Print(str)
}

func (f *Facade) Error(context string, payload map[string]interface{}, error *error, session *application_specific.Session) {
	now := time.Now()
	epoch := now.UnixNano()
	date := now.Format("2006-01-02 15:04:05.000")

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

	str := session.CorrelationId + " | " + "Date: " + strconv.FormatInt(epoch, 10) + " - " + date + " | " + "ERROR" + " | " + context + " | " + "Payload: " + payloadStr + " | " + "Error: " + errorStr

	f.printer.Print(str)
}

func (f *Facade) Clear() {
	f.printer.Clear()
}
