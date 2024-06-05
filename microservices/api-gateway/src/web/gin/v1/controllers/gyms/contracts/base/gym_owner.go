package base

import "time"

type GymOwner struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	PhoneNumber string     `json:"phoneNumber"`
	Email       string     `json:"email"`
	Restricted  bool       `json:"restricted"`
	CreatedBy   string     `json:"createdBy"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedBy   string     `json:"updatedBy"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedBy   *string    `json:"deletedBy"`
	DeletedAt   *time.Time `json:"deletedAt"`
}
