package jobs_scheduler

import (
	"github.com/robfig/cron"
	"gym-management/src/lib/primitives/application_specific"
)

type CronScheduler struct {
	scheduler *cron.Cron
}

func NewCronScheduler() *CronScheduler {
	c := cron.New()

	c.Start()

	return &CronScheduler{scheduler: c}
}

func (c *CronScheduler) ScheduleJob(job *Job) {

	err := c.scheduler.AddFunc(job.CronExpression, func() {
		err := job.Handler(application_specific.NewSession())
		if err != nil {
			// log error
		}
	})

	if err != nil {
		panic(err)
	}
}
