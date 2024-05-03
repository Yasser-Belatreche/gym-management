package get_gym_owner

import "gym-management/src/components/gyms/core/usecases/gym_owners"

type GetGymOwnerQueryResponse struct {
	gym_owners.GymOwnerToReturn
}
