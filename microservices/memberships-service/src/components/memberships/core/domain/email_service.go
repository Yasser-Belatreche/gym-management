package domain

import "gym-management-memberships/src/lib/primitives/application_specific"

type EmailService interface {
	IsUsed(email application_specific.Email, session *application_specific.Session) bool
}
