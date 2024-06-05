package customers

import (
	"gym-management-auth/src/components/auth/core/domain"
	"gym-management-auth/src/components/memberships/events"
	"gym-management-auth/src/lib/messages_broker"
	"gym-management-auth/src/lib/primitives/application_specific"
)

func BuildCustomerCreatedEventHandler(userRepository domain.UserRepository) *messages_broker.Subscriber {
	return &messages_broker.Subscriber{
		Event:     events.CustomerCreatedEventType,
		Component: "AuthManager",
		Handler: func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
			payload, err := application_specific.ParsePayload[events.CustomerCreatedEventPayload](event)
			if err != nil {
				return err
			}

			return createdEventHandler(userRepository, payload, session)
		},
	}
}

func createdEventHandler(userRepository domain.UserRepository, payload *events.CustomerCreatedEventPayload, session *application_specific.Session) *application_specific.ApplicationException {
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
		domain.RoleCustomer,
		&application_specific.UserProfile{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Phone:     payload.PhoneNumber,
			Email:     payload.Email,
		},
	)

	err = userRepository.Create(user, session)
	if err != nil {
		return err
	}

	return nil
}
