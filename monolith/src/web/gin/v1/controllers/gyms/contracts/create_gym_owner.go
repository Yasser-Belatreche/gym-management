package contracts

type CreateGymOwnerRequest struct {
	Name        string `form:"name" binding:"required"`
	Email       string `form:"email" binding:"required"`
	PhoneNumber string `form:"phoneNumber" binding:"required"`
	Password    string `form:"password" binding:"required"`
}

type CreateGymOwnerResponse struct {
	Id string `json:"id"`
}
