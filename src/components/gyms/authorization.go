package gyms

import (
	auth "gym-management/src/components/auth/core/domain"
	"gym-management/src/components/gyms/core/usecases/gym_owners/create_gym_owner"
	"gym-management/src/lib/primitives/application_specific"
)

type AuthorizationDecorator struct {
	manager Manager
}

func (d *AuthorizationDecorator) CreateGymOwner(command *create_gym_owner.CreateGymOwnerCommand) (*create_gym_owner.CreateGymOwnerCommandResponse, *application_specific.ApplicationException) {
	if command.Session.User.Role != auth.RoleAdmin {
		return nil, application_specific.NewForbiddenException("GYMS.OWNERS.CREATE_NOT_ALLOWED", "You are not allowed to create a gym owner", map[string]string{})
	}

	return d.manager.CreateGymOwner(command)
}
