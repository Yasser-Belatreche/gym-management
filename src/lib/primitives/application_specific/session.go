package application_specific

import "gym-management/src/lib/primitives/generic"

type Session struct {
	correlationId string
	user          User
}

type User struct {
	id      string
	role    string
	profile map[string]string
}

func NewSession(userId string, userRole string, profile map[string]string) *Session {
	return &Session{
		correlationId: generic.GenerateUUID(),
		user: User{
			id:      userId,
			role:    userRole,
			profile: profile,
		},
	}
}
