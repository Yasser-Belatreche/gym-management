package contracts

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
