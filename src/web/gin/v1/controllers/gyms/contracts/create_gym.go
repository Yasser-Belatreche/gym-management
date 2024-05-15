package contracts

type CreateGymRequest struct {
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

type CreateGymUrl struct {
	OwnerId string `uri:"ownerId" binding:"required"`
}

type CreateGymResponse struct {
	Id string `json:"id"`
}
