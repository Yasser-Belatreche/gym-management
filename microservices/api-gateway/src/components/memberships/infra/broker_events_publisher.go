package infra

import (
	"gym-management/src/components/memberships/core/domain/events"
	"gym-management/src/lib/messages_broker"
	"gym-management/src/lib/primitives/application_specific"
)

type BrokerEventsPublisher struct {
	Broker messages_broker.MessagesBroker
}

func (p *BrokerEventsPublisher) Publish(events []*events.MembershipEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
	for _, event := range events {
		err := p.Broker.Publish(event.DomainEvent, session)

		if err != nil {
			return err
		}
	}

	return nil
}
