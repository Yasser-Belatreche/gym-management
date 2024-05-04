package get_gym

import (
	"gym-management/src/components/gyms/core/usecases/gyms"
)

type GetGymQueryResponse struct {
	gym_owners.GymToReturn
}
