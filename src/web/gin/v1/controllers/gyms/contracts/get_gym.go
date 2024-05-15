package contracts

import "gym-management/src/web/gin/v1/controllers/gyms/contracts/base"

type GetGymUrl struct {
	Id string `uri:"gymId" binding:"required"`
}

type GetGymResponse base.Gym
