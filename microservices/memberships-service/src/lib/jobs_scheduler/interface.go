package jobs_scheduler

import "gym-management/src/lib/primitives/application_specific"

type JobsScheduler interface {
	Schedule(jobs ...*Job)
}

type Job struct {
	CronExpression string
	Handler        func(session *application_specific.Session) *application_specific.ApplicationException
}
