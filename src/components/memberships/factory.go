package memberships

import (
	"gym-management/src/components/memberships/infra"
	"gym-management/src/lib/jobs_scheduler"
	"gym-management/src/lib/messages_broker"
)

var manager Manager = nil

func NewMembershipsManager(broker messages_broker.MessagesBroker) Manager {
	if manager != nil {
		return manager
	}

	manager = &AuthorizationDecorator{
		manager: &Facade{
			CustomerRepository:   nil,
			PlanRepository:       nil,
			MembershipRepository: nil,
			EmailsService:        &infra.BrokerEmailsService{Broker: broker},
			EventsPublisher:      &infra.BrokerEventsPublisher{Broker: broker},
		},
	}

	return manager
}

func InitializeMembershipsManager(broker messages_broker.MessagesBroker, scheduler jobs_scheduler.JobsScheduler) {
	initialize(broker, scheduler, nil, nil)
}
