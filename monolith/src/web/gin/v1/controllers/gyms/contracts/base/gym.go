package base

import "time"

type Gym struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	Enabled     bool       `json:"enabled"`
	OwnerId     string     `json:"ownerId"`
	DisabledFor *string    `json:"disabledFor"`
	CreatedBy   string     `json:"createdBy"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedBy   string     `json:"updatedBy"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedBy   *string    `json:"deletedBy"`
	DeletedAt   *time.Time `json:"deletedAt"`
}
