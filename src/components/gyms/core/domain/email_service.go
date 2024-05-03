package domain

import "gym-management/src/lib/primitives/application_specific"

type EmailService interface {
	IsUsed(email application_specific.Email) bool
}
