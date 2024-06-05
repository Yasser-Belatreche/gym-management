package gym_owners

import (
	"gym-management/src/components/auth/core/domain"
	"gym-management/src/components/gyms/core/domain/events"
	"gym-management/src/lib/messages_broker"
	"gym-management/src/lib/primitives/application_specific"
)

func BuildGymOwnerDeletedEventHandler(userRepository domain.UserRepository) *messages_broker.Subscriber {
	return &messages_broker.Subscriber{
		Event:     events.GymOwnerDeletedEventType,
		Component: "AuthManager",
		Handler: func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
			payload, err := application_specific.ParsePayload[events.GymOwnerDeletedEventPayload](event)
			if err != nil {
				return err
			}

			return deletedEventHandler(userRepository, payload, session)
		},
	}
}

func deletedEventHandler(userRepository domain.UserRepository, payload *events.GymOwnerDeletedEventPayload, session *application_specific.Session) *application_specific.ApplicationException {
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
