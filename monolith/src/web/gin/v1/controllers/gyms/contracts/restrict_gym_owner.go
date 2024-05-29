package contracts

type RestrictGymOwnerUrl struct {
	Id string `uri:"ownerId" binding:"required"`
}

type RestrictGymOwnerResponse struct {
	Id string `json:"id"`
}
