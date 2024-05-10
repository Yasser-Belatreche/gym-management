package event_handlers

import (
	"gym-management/src/components/gyms/core/domain/events"
	"gym-management/src/components/memberships/core/domain"
	"gym-management/src/lib/messages_broker"
	"gym-management/src/lib/primitives/application_specific"
)

func BuildGymDisabledEventHandler(membershipRepository domain.MembershipRepository, eventsPublisher domain.EventsPublisher) *messages_broker.Subscriber {
	return &messages_broker.Subscriber{
		Event: events.GymDisabledEventType,
		Handler: func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
			payload, ok := event.Payload.(events.GymDisabledEventPayload)

			if !ok {
				return application_specific.NewDeveloperException("INVALID_EVENT_PAYLOAD_TYPE", events.GymDisabledEventType+" payload is not as expected")
			}

			return handlerGymDisabledEvent(membershipRepository, eventsPublisher, payload, session)
		},
	}
}

func handlerGymDisabledEvent(membershipRepository domain.MembershipRepository, eventsPublisher domain.EventsPublisher, payload events.GymDisabledEventPayload, session *application_specific.Session) *application_specific.ApplicationException {
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
