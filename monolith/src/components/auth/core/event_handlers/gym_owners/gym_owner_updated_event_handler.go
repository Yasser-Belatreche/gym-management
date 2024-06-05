package gym_owners

import (
	"gym-management/src/components/auth/core/domain"
	"gym-management/src/components/gyms/core/domain/events"
	"gym-management/src/lib/messages_broker"
	"gym-management/src/lib/primitives/application_specific"
)

func BuildGymOwnerUpdatedEventHandler(userRepository domain.UserRepository) *messages_broker.Subscriber {
	return &messages_broker.Subscriber{
		Event:     events.GymOwnerUpdatedEventType,
		Component: "AuthManager",
		Handler: func(event *application_specific.DomainEvent[interface{}], session *application_specific.Session) *application_specific.ApplicationException {
			payload, err := application_specific.ParsePayload[events.GymOwnerUpdatedEventPayload](event)
			if err != nil {
				return err
			}

			return updatedEventHandler(userRepository, payload, session)
		},
	}
}

func updatedEventHandler(userRepository domain.UserRepository, payload *events.GymOwnerUpdatedEventPayload, session *application_specific.Session) *application_specific.ApplicationException {
	user, err := userRepository.FindByID(payload.Id, session)
	if err != nil {
		return err
	}
	if user == nil {
		return application_specific.NewNotFoundException("USER_NOT_FOUND", "user not found", map[string]string{
			"id": payload.Id,
		})
	}

	username, err := domain.UsernameFrom(payload.Email)
	if err != nil {
		return err
	}

	user.ChangeUsernames([]domain.Username{username})
	user.SetProfile(&application_specific.UserProfile{
		FirstName:        payload.Name,
		LastName:         "",
		Phone:            payload.PhoneNumber,
		Email:            payload.Email,
		OwnedGyms:        payload.Gyms,
		EnabledOwnedGyms: payload.EnabledGyms,
	})

	if payload.NewPassword != nil {
		password, err := domain.PasswordFromPlain(*payload.NewPassword)
		if err != nil {
			return err
		}

		user.ChangePassword(password)
	}

	if payload.Restricted {
		user.Restrict()
	} else {
		user.Unrestrict()
	}

	err = userRepository.Update(user, session)
	if err != nil {
		return err
	}

	return nil
}
