package cron_jobs

import (
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/jobs_scheduler"
	"gym-management/src/lib/primitives/application_specific"
)

func BuildGenerateMonthlyBillsCronJob(membershipRepository domain.MembershipRepository, eventsPublisher domain.EventsPublisher) *jobs_scheduler.Job {
	const firstDayOfEveryMonth = "0 0 0 1 * *"

	return &jobs_scheduler.Job{
		CronExpression: firstDayOfEveryMonth,
		Handler: func(session *application_specific.Session) *application_specific.ApplicationException {
			return generateBillsHandler(membershipRepository, eventsPublisher, session)
		},
	}
}

func generateBillsHandler(membershipRepository domain.MembershipRepository, eventsPublisher domain.EventsPublisher, session *application_specific.Session) *application_specific.ApplicationException {
	memberships, err := membershipRepository.FindEnabledMemberships(session)
	if err != nil {
		return err
	}

	for _, membership := range memberships {
		err = membership.GenerateMonthlyBill()
		if err != nil {
			return err
		}

		err = membershipRepository.Update(membership, session)
		if err != nil {
			return err
		}

		err = eventsPublisher.Publish(membership.PullEvents(), session)
		if err != nil {
			return err
		}
	}

	return nil
}
