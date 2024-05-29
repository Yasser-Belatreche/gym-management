package contracts

type UpdateGymOwnerRequest struct {
	Name        string  `form:"name" binding:"required"`
	Email       string  `form:"email" binding:"required"`
	PhoneNumber string  `form:"phoneNumber" binding:"required"`
	NewPassword *string `form:"password" binding:"-"`
}

type UpdateGymOwnerUrl struct {
	Id string `uri:"ownerId" binding:"required"`
}

type UpdateGymOwnerResponse struct {
	Id string `json:"id"`
}
