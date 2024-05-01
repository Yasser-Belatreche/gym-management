package application_specific

import "gym-management/src/lib/primitives/generic"

type Session struct {
	correlationId string
}

type UserSession struct {
	Session
	user User
}

type User struct {
	id      string
	role    string
	profile map[string]string
}

func NewSession() *Session {
	return &Session{
		correlationId: generic.GenerateUUID(),
	}
}

func NewUserSession(userId string, userRole string, profile map[string]string) *UserSession {
	return &UserSession{
		Session: Session{
			correlationId: generic.GenerateUUID(),
		},
		user: User{
			id:      userId,
			role:    userRole,
			profile: profile,
		},
	}
}
