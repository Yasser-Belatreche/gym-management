package contracts

type UpdateGymRequest struct {
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

type UpdateGymUrl struct {
	OwnerId string `uri:"ownerId" binding:"required"`
	GymId   string `uri:"gymId" binding:"required"`
}

type UpdateGymResponse struct {
	Id string `json:"id"`
}
