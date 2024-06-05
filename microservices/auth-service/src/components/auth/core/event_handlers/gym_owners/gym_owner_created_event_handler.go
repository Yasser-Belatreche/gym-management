package gym_owners

import (
	"gym-management-auth/src/components/auth/core/domain"
	"gym-management-auth/src/components/gyms/events"
	"gym-management-auth/src/lib/messages_broker"
	"gym-management-auth/src/lib/primitives/application_specific"
)

func BuildGymOwnerCreatedEventHandler(userRepository domain.UserRepository) *messages_broker.Subscriber {
	return &messages_broker.Subscriber{
		Event:     events.GymOwnerCreatedEventType,
		Component: "AuthManager",
		Handler: func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
			payload, err := application_specific.ParsePayload[events.GymOwnerCreatedEventPayload](event)
			if err != nil {
				return err
			}

			return createdEventHandler(userRepository, *payload, session)
		},
	}
}

func createdEventHandler(userRepository domain.UserRepository, payload events.GymOwnerCreatedEventPayload, session *application_specific.Session) *application_specific.ApplicationException {
	username, err := domain.UsernameFrom(payload.Email)
	if err != nil {
		return err
	}
	password, err := domain.PasswordFromPlain(payload.Password)
	if err != nil {
		return err
	}

	user := domain.CreateUser(
		payload.Id,
		[]domain.Username{username},
		password,
		domain.RoleGymOwner,
		&application_specific.UserProfile{
			FirstName:        payload.Name,
			LastName:         "",
			Phone:            payload.PhoneNumber,
			Email:            payload.Email,
			OwnedGyms:        payload.Gyms,
			EnabledOwnedGyms: payload.Gyms,
		},
	)

	err = userRepository.Create(user, session)
	if err != nil {
		return err
	}

	return nil
}
