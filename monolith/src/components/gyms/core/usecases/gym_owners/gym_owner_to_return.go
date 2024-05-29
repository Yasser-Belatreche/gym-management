package gym_owners

import "time"

type GymOwnerToReturn struct {
	Id          string
	Name        string
	PhoneNumber string
	Email       string
	Restricted  bool
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedBy   string
	UpdatedAt   time.Time
	DeletedBy   *string
	DeletedAt   *time.Time
}
