package contracts

import "gym-management-gyms/src/web/gin/v1/controllers/gyms/contracts/base"

type GetGymOwnerUrl struct {
	Id string `uri:"ownerId" binding:"required"`
}

type GetGymOwnerResponse base.GymOwner
