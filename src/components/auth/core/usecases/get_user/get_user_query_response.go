package get_user

import "time"

type GetUserQueryResponse struct {
	Id        string
	Role      string
	Email     string
	Usernames []string
	FirstName string
	LastName  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	LastLogin *time.Time
	CreatedBy string
	UpdatedBy string
	DeletedAt *time.Time
	DeletedBy *string
}
