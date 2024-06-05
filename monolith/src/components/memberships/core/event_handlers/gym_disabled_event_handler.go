package event_handlers

import (
	"gym-management/src/components/gyms/core/domain/events"
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/messages_broker"
	"gym-management/src/lib/primitives/application_specific"
)

func BuildGymDisabledEventHandler(membershipRepository domain.MembershipRepository, eventsPublisher domain.EventsPublisher) *messages_broker.Subscriber {
	return &messages_broker.Subscriber{
		Event:     events.GymDisabledEventType,
		Component: "MembershipsManager",
		Handler: func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
			payload, err := application_specific.ParsePayload[events.GymDisabledEventPayload](event)
			if err != nil {
				return err
			}

			return handlerGymDisabledEvent(membershipRepository, eventsPublisher, payload, session)
		},
	}
}

func handlerGymDisabledEvent(membershipRepository domain.MembershipRepository, eventsPublisher domain.EventsPublisher, payload *events.GymDisabledEventPayload, session *application_specific.Session) *application_specific.ApplicationException {
	gymId := payload.Id
	disabledBy := payload.DisabledBy

	memberships, err := membershipRepository.FindEnabledMembershipsByGymID(gymId, session)
	if err != nil {
		return err
	}

	for _, membership := range memberships {
		err = membership.GymDisabled(disabledBy)
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
