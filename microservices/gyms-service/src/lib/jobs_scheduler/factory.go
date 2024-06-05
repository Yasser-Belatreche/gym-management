package jobs_scheduler

var scheduler JobsScheduler = nil

func NewJobsScheduler() JobsScheduler {
	if scheduler == nil {
		scheduler = NewCronScheduler()
	}

	return scheduler
}
