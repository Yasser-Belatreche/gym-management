package lib

import (
	"gorm.io/gorm"
	"gym-management-gyms/src/lib/jobs_scheduler"
	"gym-management-gyms/src/lib/logger"
	"gym-management-gyms/src/lib/messages_broker"
	"gym-management-gyms/src/lib/persistence"
	"gym-management-gyms/src/lib/primitives/application_specific"
)

func InitializeLib() {
	persistence.InitializePersistence()
}

func MessagesBroker() messages_broker.MessagesBroker {
	return messages_broker.NewMessagesBroker()
}

func JobsScheduler() jobs_scheduler.JobsScheduler {
	return jobs_scheduler.NewJobsScheduler()
}

func Persistence() persistence.Persistence {
	return persistence.NewPersistence()
}

func Logger() logger.Logger {
	return logger.NewLogger()
}

func GormDB(session interface{}) *gorm.DB {
	switch e := session.(type) {
	case *application_specific.Session:
		return persistence.NewGormPersistence().GetClient(e)
	case *application_specific.UserSession:
		return persistence.NewGormPersistence().GetClient(e.Session)
	case application_specific.Session:
		return persistence.NewGormPersistence().GetClient(&e)
	case application_specific.UserSession:
		return persistence.NewGormPersistence().GetClient(e.Session)
	default:
		panic("Invalid session type")
	}
}
