package base

import "time"

type Plan struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	Featured       bool       `json:"featured"`
	SessionPerWeek int        `json:"sessionPerWeek"`
	WithCoach      bool       `json:"withCoach"`
	MonthlyPrice   float64    `json:"monthlyPrice"`
	GymId          string     `json:"gymId"`
	CreatedBy      string     `json:"createdBy"`
	UpdatedBy      string     `json:"updatedBy"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	DeletedBy      *string    `json:"deletedBy"`
	DeletedAt      *time.Time `json:"deletedAt"`
}
