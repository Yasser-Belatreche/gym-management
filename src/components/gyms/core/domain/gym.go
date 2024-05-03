package domain

import "time"

type Gym struct {
	id        string
	name      string
	address   string
	enabled   bool
	createdBy string
	updatedBy string
	deletedAt *time.Time
	deleteBy  *string
}
