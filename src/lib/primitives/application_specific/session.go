package application_specific

import "gym-management/src/lib/primitives/generic"

type Session struct {
	CorrelationId string
}

type UserSession struct {
	*Session
	User *User
}

type User struct {
	Id      string
	Role    string
	Profile *UserProfile
}

type UserProfile struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string

	OwnedGyms        []string // in case the user is a gym owner
	EnabledOwnedGyms []string // in case the user is a gym owner
}

func NewSession() *Session {
	return &Session{
		CorrelationId: generic.GenerateUUID(),
	}
}

func NewUserSession(userId string, userRole string, profile *UserProfile) *UserSession {
	return &UserSession{
		Session: &Session{
			CorrelationId: generic.GenerateUUID(),
		},
		User: &User{
			Id:      userId,
			Role:    userRole,
			Profile: profile,
		},
	}
}
