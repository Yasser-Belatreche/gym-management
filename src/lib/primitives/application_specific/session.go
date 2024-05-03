package application_specific

import "gym-management/src/lib/primitives/generic"

type Session struct {
	CorrelationId string
}

type UserSession struct {
	Session
	User *User
}

type User struct {
	Id      string
	Role    string
	Profile map[string]string
}

func NewSession() *Session {
	return &Session{
		CorrelationId: generic.GenerateUUID(),
	}
}

func NewUserSession(userId string, userRole string, profile map[string]string) *UserSession {
	return &UserSession{
		Session: Session{
			CorrelationId: generic.GenerateUUID(),
		},
		User: &User{
			Id:      userId,
			Role:    userRole,
			Profile: profile,
		},
	}
}
