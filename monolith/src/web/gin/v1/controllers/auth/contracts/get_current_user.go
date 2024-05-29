package contracts

import "time"

type GetCurrentUserRequest struct{}

type GetCurrentUserResponse struct {
	Id        string     `json:"id"`
	Role      string     `json:"role"`
	Usernames []string   `json:"usernames"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	LastLogin *time.Time `json:"lastLogin"`
}
