package jobs_scheduler

import (
	"github.com/robfig/cron"
	"gym-management-memberships/src/lib/persistence"
	"gym-management-memberships/src/lib/primitives/application_specific"
)

type CronScheduler struct {
	scheduler *cron.Cron
}

func NewCronScheduler() *CronScheduler {
	c := cron.New()

	c.Start()

	return &CronScheduler{scheduler: c}
}

func (c *CronScheduler) Schedule(jobs ...*Job) {
	for _, job := range jobs {
		err := c.scheduler.AddFunc(job.CronExpression, func() {
			session := application_specific.NewSession()

			err := persistence.NewPersistence().WithTransaction(session, func() *application_specific.ApplicationException {
				err := job.Handler(session)
				if err != nil {
					// log error
				}

				return err
			})

			if err != nil {
				panic(err)
			}
		})

		if err != nil {
			panic(err)
		}
	}
}
