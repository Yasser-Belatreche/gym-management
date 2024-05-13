package persistence

import "gym-management/src/lib/primitives/application_specific"

type Persistence interface {
	Connect() *application_specific.ApplicationException

	Disconnect() *application_specific.ApplicationException

	WithTransaction(session *application_specific.Session, method func() *application_specific.ApplicationException) *application_specific.ApplicationException

	HealthCheck() *PersistenceHealth
}

type PersistenceHealth struct {
	Status   string
	Provider string
	Message  string
}
