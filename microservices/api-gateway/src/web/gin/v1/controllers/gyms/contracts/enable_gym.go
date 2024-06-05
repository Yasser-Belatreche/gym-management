package contracts

type EnableGymUrl struct {
	OwnerId string `uri:"ownerId" binding:"required"`
	GymId   string `uri:"gymId" binding:"required"`
}

type EnableGymResponse struct {
	Id string `json:"id"`
}
