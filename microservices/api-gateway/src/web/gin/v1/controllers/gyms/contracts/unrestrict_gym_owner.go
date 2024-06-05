package contracts

type UnrestrictGymOwnerUrl struct {
	Id string `uri:"ownerId" binding:"required"`
}

type UnrestrictGymOwnerResponse struct {
	Id string `json:"id"`
}
