package contracts

type DisableGymUrl struct {
	OwnerId string `uri:"ownerId" binding:"required"`
	GymId   string `uri:"gymId" binding:"required"`
}

type DisableGymResponse struct {
	Id string `json:"id"`
}
