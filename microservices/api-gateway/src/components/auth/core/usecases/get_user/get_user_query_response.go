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
	LastLogin *time.Time
	DeletedAt *time.Time
}
