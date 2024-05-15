package memberships

import (
	"gym-management/src/components/memberships/core/cron_jobs"
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/jobs_scheduler"
	"gym-management/src/lib/messages_broker"
)

func initialize(
	broker messages_broker.MessagesBroker,
	scheduler jobs_scheduler.JobsScheduler,
	eventsPublisher domain.EventsPublisher,
	membershipRepository domain.MembershipRepository,
) {
	//broker.Subscribe(
	//	event_handlers.BuildGymDisabledEventHandler(
	//		membershipRepository,
	//		eventsPublisher,
	//	),
	//)

	scheduler.Schedule(
		cron_jobs.BuildGenerateMonthlyBillsCronJob(membershipRepository, eventsPublisher),
	)
}
