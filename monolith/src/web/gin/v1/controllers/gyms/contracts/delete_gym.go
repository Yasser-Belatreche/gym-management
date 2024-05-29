package contracts

type DeleteGymUrl struct {
	OwnerId string `uri:"ownerId" binding:"required"`
	GymId   string `uri:"gymId" binding:"required"`
}

type DeleteGymResponse struct {
	Id string `json:"id"`
}
