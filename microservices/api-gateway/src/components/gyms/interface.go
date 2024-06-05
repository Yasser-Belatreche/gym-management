package gyms

import (
	"gym-management/src/components/gyms/core/usecases/gym_owners/create_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/delete_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/get_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/get_gym_owners"
	"gym-management/src/components/gyms/core/usecases/gym_owners/restrict_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/unrestrict_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/update_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gyms/create_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/delete_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/disable_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/enable_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/get_gym"
	"gym-management/src/components/gyms/core/usecases/gyms/get_gyms"
	"gym-management/src/components/gyms/core/usecases/gyms/update_gym"
	"gym-management/src/lib/primitives/application_specific"
)

type Manager interface {
	GetGymOwner(query *get_gym_owner.GetGymOwnerQuery) (*get_gym_owner.GetGymOwnerQueryResponse, *application_specific.ApplicationException)

	GetGymOwners(query *get_gym_owners.GetGymOwnersQuery) (*get_gym_owners.GetGymOwnersQueryResponse, *application_specific.ApplicationException)

	CreateGymOwner(command *create_gym_owner.CreateGymOwnerCommand) (*create_gym_owner.CreateGymOwnerCommandResponse, *application_specific.ApplicationException)

	UpdateGymOwner(command *update_gym_owner.UpdateGymOwnerCommand) (*update_gym_owner.UpdateGymOwnerCommandResponse, *application_specific.ApplicationException)

	DeleteGymOwner(command *delete_gym_owner.DeleteGymOwnerCommand) (*delete_gym_owner.DeleteGymOwnerCommandResponse, *application_specific.ApplicationException)

	RestrictGymOwner(command *restrict_gym_owner.RestrictGymOwnerCommand) (*restrict_gym_owner.RestrictGymOwnerCommandResponse, *application_specific.ApplicationException)

	UnrestrictGymOwner(command *unrestrict_gym_owner.UnrestrictGymOwnerCommand) (*unrestrict_gym_owner.UnrestrictGymOwnerCommandResponse, *application_specific.ApplicationException)

	GetGym(query *get_gym.GetGymQuery) (*get_gym.GetGymQueryResponse, *application_specific.ApplicationException)

	GetGyms(query *get_gyms.GetGymsQuery) (*get_gyms.GetGymsQueryResponse, *application_specific.ApplicationException)

	CreateGym(command *create_gym.CreateGymCommand) (*create_gym.CreateGymCommandResponse, *application_specific.ApplicationException)

	UpdateGym(command *update_gym.UpdateGymCommand) (*update_gym.UpdateGymCommandResponse, *application_specific.ApplicationException)

	DeleteGym(command *delete_gym.DeleteGymCommand) (*delete_gym.DeleteGymCommandResponse, *application_specific.ApplicationException)

	EnableGym(command *enable_gym.EnableGymCommand) (*enable_gym.EnableGymCommandResponse, *application_specific.ApplicationException)

	DisableGym(command *disable_gym.DisableGymCommand) (*disable_gym.DisableGymCommandResponse, *application_specific.ApplicationException)
}
