package lib

import (
	"gorm.io/gorm"
	"gym-management/src/lib/jobs_scheduler"
	"gym-management/src/lib/messages_broker"
	"gym-management/src/lib/persistence"
	"gym-management/src/lib/primitives/application_specific"
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
