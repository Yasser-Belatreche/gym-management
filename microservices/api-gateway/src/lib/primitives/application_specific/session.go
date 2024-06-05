package application_specific

import (
	"encoding/base64"
	"encoding/json"
)

type Session struct {
	CorrelationId string `json:"correlationId"`
}

func (s *Session) ToBase64() (string, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", NewUnknownException("ERROR_MARSHALLING_SESSION", err.Error(), nil)
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
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

func (s *UserSession) ToBase64() (string, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", NewUnknownException("ERROR_MARSHALLING_USER_SESSION", err.Error(), nil)
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}
