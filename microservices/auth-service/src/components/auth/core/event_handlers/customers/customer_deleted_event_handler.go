package customers

import (
	"gym-management-auth/src/components/auth/core/domain"
	"gym-management-auth/src/components/memberships/events"
	"gym-management-auth/src/lib/messages_broker"
	"gym-management-auth/src/lib/primitives/application_specific"
)

func BuildCustomerDeletedEventHandler(userRepository domain.UserRepository) *messages_broker.Subscriber {
	return &messages_broker.Subscriber{
		Event:     events.CustomerDeletedEventType,
		Component: "AuthManager",
		Handler: func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
			payload, err := application_specific.ParsePayload[events.CustomerDeletedEventPayload](event)
			if err != nil {
				return err
			}

			return deletedEventHandler(userRepository, payload, session)
		},
	}
}

func deletedEventHandler(userRepository domain.UserRepository, payload *events.CustomerDeletedEventPayload, session *application_specific.Session) *application_specific.ApplicationException {
	user, err := userRepository.FindByID(payload.Id, session)
	if err != nil {
		return err
	}
	if user == nil {
		return application_specific.NewNotFoundException("USER_NOT_FOUND", "user not found", map[string]string{
			"id": payload.Id,
		})
	}

	user.Delete(payload.DeletedBy)

	err = userRepository.Update(user, session)
	if err != nil {
		return err
	}

	return nil
}
