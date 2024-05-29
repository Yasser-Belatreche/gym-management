package gyms

import (
	"gym-management/src/components/gyms/infra"
	"gym-management/src/lib/messages_broker"
)

var manager Manager = nil

func NewGymsManager(broker messages_broker.MessagesBroker) Manager {
	if manager == nil {
		manager = &AuthorizationDecorator{
			manager: &Facade{
				EmailService:       &infra.BrokerEmailsService{Broker: broker},
				EventsPublisher:    &infra.BrokerEventsPublisher{Broker: broker},
				GymOwnerRepository: &infra.GormGymOwnerRepository{},
			},
		}
	}

	return manager
}

func InitializeGymsManager() {
	initialize()
}
