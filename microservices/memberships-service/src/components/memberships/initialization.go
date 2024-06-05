package memberships

import (
	"gym-management-memberships/src/components/memberships/core/cron_jobs"
	"gym-management-memberships/src/components/memberships/core/domain"
	"gym-management-memberships/src/components/memberships/core/event_handlers"
	"gym-management-memberships/src/lib/jobs_scheduler"
	"gym-management-memberships/src/lib/messages_broker"
)

func initialize(
	broker messages_broker.MessagesBroker,
	scheduler jobs_scheduler.JobsScheduler,
	eventsPublisher domain.EventsPublisher,
	membershipRepository domain.MembershipRepository,
) {
	broker.Subscribe(
		event_handlers.BuildGymDisabledEventHandler(
			membershipRepository,
			eventsPublisher,
		),
	)

	scheduler.Schedule(
		cron_jobs.BuildGenerateMonthlyBillsCronJob(membershipRepository, eventsPublisher),
	)
}
