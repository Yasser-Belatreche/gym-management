package application_specific

import "gym-management/src/lib/primitives/generic"

type Session struct {
	CorrelationId string
}

func NewSession() *Session {
	return &Session{
		CorrelationId: generic.GenerateUUID(),
	}
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
