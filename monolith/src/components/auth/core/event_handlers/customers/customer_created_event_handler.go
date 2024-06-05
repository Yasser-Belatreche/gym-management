package customers

import (
	"github.com/mitchellh/mapstructure"
	"gym-management/src/components/auth/core/domain"
	"gym-management/src/components/memberships/core/domain/events"
	"gym-management/src/lib/messages_broker"
	"gym-management/src/lib/primitives/application_specific"
)

func BuildCustomerCreatedEventHandler(userRepository domain.UserRepository) *messages_broker.Subscriber {
	return &messages_broker.Subscriber{
		Event:     events.CustomerCreatedEventType,
		Component: "AuthManager",
		Handler: func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
			switch payload := event.Payload.(type) {
			case events.CustomerCreatedEventPayload:
				return createdEventHandler(userRepository, payload, session)
			case *events.CustomerCreatedEventPayload:
				return createdEventHandler(userRepository, *payload, session)
			case map[string]interface{}:
				var decodedPayload events.CustomerCreatedEventPayload

				err := mapstructure.Decode(payload, &decodedPayload)
				if err != nil {
					return application_specific.NewDeveloperException("INVALID_EVENT_PAYLOAD_TYPE", events.CustomerCreatedEventType+" payload is not as expected")
				}

				return createdEventHandler(userRepository, decodedPayload, session)
			default:
				return application_specific.NewDeveloperException("INVALID_EVENT_PAYLOAD_TYPE", events.CustomerCreatedEventType+" payload is not as expected")
			}
		},
	}
}

func createdEventHandler(userRepository domain.UserRepository, payload events.CustomerCreatedEventPayload, session *application_specific.Session) *application_specific.ApplicationException {
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
