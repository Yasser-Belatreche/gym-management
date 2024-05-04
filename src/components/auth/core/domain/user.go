package domain

import (
	"gym-management/src/lib/primitives/application_specific"
	"time"
)

type User struct {
	id         string
	usernames  []Username
	password   Password
	role       string
	profile    *application_specific.UserProfile
	restricted bool
	createdBy  string
	updatedBy  string
	deletedAt  *time.Time
	deletedBy  *string
}

func NewUser(id string, usernames []Username, password Password, role string, profile *application_specific.UserProfile, createdBy string) *User {
	return &User{
		id:         id,
		usernames:  usernames,
		password:   password,
		role:       role,
		profile:    profile,
		restricted: false,
		deletedBy:  nil,
		updatedBy:  createdBy,
		deletedAt:  nil,
		createdBy:  createdBy,
	}
}
