package gyms

import (
	"gym-management/src/components/gyms/core/usecases/gym_owners/create_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/delete_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/restrict_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/unrestrict_gym_owner"
	"gym-management/src/components/gyms/core/usecases/gym_owners/update_gym_owner"
	"gym-management/src/lib/primitives/application_specific"
)

type Manager interface {
	CreateGymOwner(command *create_gym_owner.CreateGymOwnerCommand) (*create_gym_owner.CreateGymOwnerCommandResponse, *application_specific.ApplicationException)

	UpdateGymOwner(command *update_gym_owner.UpdateGymOwnerCommand) (*update_gym_owner.UpdateGymOwnerCommandResponse, *application_specific.ApplicationException)

	DeleteGymOwner(command *delete_gym_owner.DeleteGymOwnerCommand) (*delete_gym_owner.DeleteGymOwnerCommandResponse, *application_specific.ApplicationException)

	RestrictGymOwner(command *restrict_gym_owner.RestrictGymOwnerCommand) (*restrict_gym_owner.RestrictGymOwnerCommandResponse, *application_specific.ApplicationException)

	UnrestrictGymOwner(command *unrestrict_gym_owner.UnrestrictGymOwnerCommand) (*unrestrict_gym_owner.UnrestrictGymOwnerCommandResponse, *application_specific.ApplicationException)
}
