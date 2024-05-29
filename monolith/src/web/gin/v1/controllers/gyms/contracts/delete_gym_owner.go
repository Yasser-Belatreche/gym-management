package contracts

type DeleteGymOwnerUrl struct {
	Id string `uri:"ownerId" binding:"required"`
}

type DeleteGymOwnerResponse struct {
	Id string `json:"id"`
}
