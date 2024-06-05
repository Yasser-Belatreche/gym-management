package logger

import "gym-management-gyms/src/lib/primitives/application_specific"

type Logger interface {
	Info(
		msg string,
		payload map[string]interface{},
		session *application_specific.Session,
	)

	Warn(
		msg string,
		payload map[string]interface{},
		error *error,
		session *application_specific.Session,
	)

	Error(
		context string,
		payload map[string]interface{},
		error *error,
		session *application_specific.Session,
	)

	Clear()
}
