package lib

import (
	"gym-management/src/lib/jobs_scheduler"
	"gym-management/src/lib/messages_broker"
)

func MessagesBroker() messages_broker.MessagesBroker {
	return messages_broker.NewMessagesBroker()
}

func JobsScheduler() jobs_scheduler.JobsScheduler {
	return jobs_scheduler.NewJobsScheduler()
}
