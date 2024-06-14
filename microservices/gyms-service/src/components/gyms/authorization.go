package gyms

import (
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/create_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/delete_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/get_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/get_gym_owners"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/restrict_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/unrestrict_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gym_owners/update_gym_owner"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/create_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/delete_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/disable_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/enable_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/get_gym"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/get_gyms"
	"gym-management-gyms/src/components/gyms/core/usecases/gyms/update_gym"
	"gym-management-gyms/src/lib/primitives/application_specific"
)

type AuthorizationDecorator struct {
	manager Manager
}

func (d *AuthorizationDecorator) GetGymOwner(query *get_gym_owner.GetGymOwnerQuery) (*get_gym_owner.GetGymOwnerQueryResponse, *application_specific.ApplicationException) {
	owner, err := d.manager.GetGymOwner(query)
	if err != nil {
		return owner, err
	}

	if query.Session.User.Role == application_specific.RoleAdmin {
		return owner, err
	}

	if query.Session.User.Role == application_specific.RoleGymOwner {
		if query.Session.User.Id == owner.Id {
			return owner, err
		}
	}

	return nil, application_specific.NewForbiddenException("GYMS.OWNERS.GET_NOT_ALLOWED", "You are not allowed to get this gym owner", map[string]string{
		"id": owner.Id,
	})
}

func (d *AuthorizationDecorator) GetGymOwners(query *get_gym_owners.GetGymOwnersQuery) (*get_gym_owners.GetGymOwnersQueryResponse, *application_specific.ApplicationException) {
	if query.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.GetGymOwners(query)
	}

	if query.Session.User.Role == application_specific.RoleGymOwner {
		query.Id = make([]string, 0)
		query.Id = append(query.Id, query.Session.User.Id)

		return d.manager.GetGymOwners(query)
	}

	return nil, application_specific.NewForbiddenException("GYMS.OWNERS.GET_NOT_ALLOWED", "You are not allowed to get gym owners", map[string]string{})
}

func (d *AuthorizationDecorator) CreateGymOwner(command *create_gym_owner.CreateGymOwnerCommand) (*create_gym_owner.CreateGymOwnerCommandResponse, *application_specific.ApplicationException) {
	if command.Session.User.Role != application_specific.RoleAdmin {
		return nil, application_specific.NewForbiddenException("GYMS.OWNERS.CREATE_NOT_ALLOWED", "You are not allowed to create a gym owner", map[string]string{})
	}

	return d.manager.CreateGymOwner(command)
}

func (d *AuthorizationDecorator) UpdateGymOwner(command *update_gym_owner.UpdateGymOwnerCommand) (*update_gym_owner.UpdateGymOwnerCommandResponse, *application_specific.ApplicationException) {
	owner, err := d.manager.GetGymOwner(&get_gym_owner.GetGymOwnerQuery{
		Id:      command.Id,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if command.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.UpdateGymOwner(command)
	}

	if command.Session.User.Role == application_specific.RoleGymOwner {
		if command.Session.User.Id == owner.Id {
			return d.manager.UpdateGymOwner(command)
		}
	}

	return nil, application_specific.NewForbiddenException("GYMS.OWNERS.UPDATE_NOT_ALLOWED", "You are not allowed to update this gym owner", map[string]string{
		"id": owner.Id,
	})
}

func (d *AuthorizationDecorator) DeleteGymOwner(command *delete_gym_owner.DeleteGymOwnerCommand) (*delete_gym_owner.DeleteGymOwnerCommandResponse, *application_specific.ApplicationException) {
	if command.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.DeleteGymOwner(command)
	}

	return nil, application_specific.NewForbiddenException("GYMS.OWNERS.DELETE_NOT_ALLOWED", "You are not allowed to delete this gym owner", map[string]string{
		"id": command.Id,
	})
}

func (d *AuthorizationDecorator) RestrictGymOwner(command *restrict_gym_owner.RestrictGymOwnerCommand) (*restrict_gym_owner.RestrictGymOwnerCommandResponse, *application_specific.ApplicationException) {
	if command.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.RestrictGymOwner(command)
	}

	return nil, application_specific.NewForbiddenException("GYMS.OWNERS.RESTRICT_NOT_ALLOWED", "You are not allowed to restrict this gym owner", map[string]string{
		"id": command.Id,
	})
}

func (d *AuthorizationDecorator) UnrestrictGymOwner(command *unrestrict_gym_owner.UnrestrictGymOwnerCommand) (*unrestrict_gym_owner.UnrestrictGymOwnerCommandResponse, *application_specific.ApplicationException) {
	if command.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.UnrestrictGymOwner(command)
	}

	return nil, application_specific.NewForbiddenException("GYMS.OWNERS.UNRESTRICT_NOT_ALLOWED", "You are not allowed to unrestrict this gym owner", map[string]string{
		"id": command.Id,
	})
}

func (d *AuthorizationDecorator) GetGym(query *get_gym.GetGymQuery) (*get_gym.GetGymQueryResponse, *application_specific.ApplicationException) {
	gym, err := d.manager.GetGym(query)
	if err != nil {
		return gym, err
	}

	if query.Session.User.Role == application_specific.RoleAdmin {
		return gym, err
	}

	if query.Session.User.Role == application_specific.RoleGymOwner {
		if query.Session.User.Id == gym.OwnerId {
			return gym, err
		}
	}

	return nil, application_specific.NewForbiddenException("GYMS.GET_NOT_ALLOWED", "You are not allowed to get this gym", map[string]string{
		"id": gym.Id,
	})
}

func (d *AuthorizationDecorator) GetGyms(query *get_gyms.GetGymsQuery) (*get_gyms.GetGymsQueryResponse, *application_specific.ApplicationException) {
	if query.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.GetGyms(query)
	}

	if query.Session.User.Role == application_specific.RoleGymOwner {
		query.OwnerId = make([]string, 0)
		query.OwnerId = append(query.OwnerId, query.Session.User.Id)

		return d.manager.GetGyms(query)
	}

	return nil, application_specific.NewForbiddenException("GYMS.GET_NOT_ALLOWED", "You are not allowed to get gyms", map[string]string{})
}

func (d *AuthorizationDecorator) CreateGym(command *create_gym.CreateGymCommand) (*create_gym.CreateGymCommandResponse, *application_specific.ApplicationException) {
	if command.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.CreateGym(command)
	}

	if command.Session.User.Role == application_specific.RoleGymOwner {
		command.OwnerId = command.Session.User.Id

		return d.manager.CreateGym(command)
	}

	return nil, application_specific.NewForbiddenException("GYMS.CREATE_NOT_ALLOWED", "You are not allowed to create a gym", map[string]string{})
}

func (d *AuthorizationDecorator) UpdateGym(command *update_gym.UpdateGymCommand) (*update_gym.UpdateGymCommandResponse, *application_specific.ApplicationException) {
	gym, err := d.manager.GetGym(&get_gym.GetGymQuery{
		Id:      command.GymId,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if command.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.UpdateGym(command)
	}

	if command.Session.User.Role == application_specific.RoleGymOwner {
		command.OwnerId = command.Session.User.Id

		if command.Session.User.Id == gym.OwnerId {
			return d.manager.UpdateGym(command)
		}
	}

	return nil, application_specific.NewForbiddenException("GYMS.UPDATE_NOT_ALLOWED", "You are not allowed to update this gym", map[string]string{
		"id": gym.Id,
	})
}

func (d *AuthorizationDecorator) DeleteGym(command *delete_gym.DeleteGymCommand) (*delete_gym.DeleteGymCommandResponse, *application_specific.ApplicationException) {
	gym, err := d.manager.GetGym(&get_gym.GetGymQuery{
		Id:      command.GymId,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if command.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.DeleteGym(command)
	}

	if command.Session.User.Role == application_specific.RoleGymOwner {
		command.OwnerId = command.Session.User.Id

		if command.Session.User.Id == gym.OwnerId {
			return d.manager.DeleteGym(command)
		}
	}

	return nil, application_specific.NewForbiddenException("GYMS.DELETE_NOT_ALLOWED", "You are not allowed to delete this gym", map[string]string{
		"id": gym.Id,
	})
}

func (d *AuthorizationDecorator) EnableGym(command *enable_gym.EnableGymCommand) (*enable_gym.EnableGymCommandResponse, *application_specific.ApplicationException) {
	gym, err := d.manager.GetGym(&get_gym.GetGymQuery{
		Id:      command.GymId,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if command.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.EnableGym(command)
	}

	if command.Session.User.Role == application_specific.RoleGymOwner {
		command.OwnerId = command.Session.User.Id

		if command.Session.User.Id == gym.OwnerId {
			return d.manager.EnableGym(command)
		}
	}

	return nil, application_specific.NewForbiddenException("GYMS.ENABLE_NOT_ALLOWED", "You are not allowed to enable this gym", map[string]string{
		"id": gym.Id,
	})
}

func (d *AuthorizationDecorator) DisableGym(command *disable_gym.DisableGymCommand) (*disable_gym.DisableGymCommandResponse, *application_specific.ApplicationException) {
	gym, err := d.manager.GetGym(&get_gym.GetGymQuery{
		Id:      command.GymId,
		Session: command.Session,
	})
	if err != nil {
		return nil, err
	}

	if command.Session.User.Role == application_specific.RoleAdmin {
		return d.manager.DisableGym(command)
	}

	if command.Session.User.Role == application_specific.RoleGymOwner {
		command.OwnerId = command.Session.User.Id

		if command.Session.User.Id == gym.OwnerId {
			return d.manager.DisableGym(command)
		}
	}

	return nil, application_specific.NewForbiddenException("GYMS.DISABLE_NOT_ALLOWED", "You are not allowed to disable this gym", map[string]string{
		"id": gym.Id,
	})
}
