package application_specific

import (
	"gym-management-auth/src/lib/primitives/generic"
)

type Session struct {
	CorrelationId string `json:"correlationId"`
}

func NewSession() *Session {
	return &Session{
		CorrelationId: generic.GenerateUUID(),
	}
}

func NewSessionWithCorrelationId(correlationId string) *Session {
	return &Session{
		CorrelationId: correlationId,
	}
}

type UserSession struct {
	*Session
	User *User `json:"user"`
}

type User struct {
	Id      string       `json:"id"`
	Role    string       `json:"role"`
	Profile *UserProfile `json:"profile"`
}

type UserProfile struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`

	OwnedGyms        []string `json:"ownedGyms"`        // in case the user is a gym owner
	EnabledOwnedGyms []string `json:"enabledOwnedGyms"` // in case the user is a gym owner
}

func NewUserSession(userId string, userRole string, profile *UserProfile, session *Session) *UserSession {
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

func (s *UserSession) UserId() string {
	return s.User.Id
}

func (s *UserSession) RoleIsOneOf(roles ...string) bool {
	for _, role := range roles {
		if s.User.Role == role {
			return true
		}
	}

	return false
}

func (s *UserSession) IsOwnerOfEnabledGym(gymId string) bool {
	if s.User.Profile.EnabledOwnedGyms == nil {
		return false
	}

	for _, gym := range s.User.Profile.EnabledOwnedGyms {
		if gym == gymId {
			return true
		}
	}

	return false
}
