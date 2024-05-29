package gyms

import "time"

type GymToReturn struct {
	Id          string
	Name        string
	Address     string
	Enabled     bool
	OwnerId     string
	DisabledFor *string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedBy   string
	UpdatedAt   time.Time
	DeletedBy   *string
	DeletedAt   *time.Time
}
