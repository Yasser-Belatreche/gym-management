package domain

import "gym-management-memberships/src/lib/primitives/application_specific"

type PlanRepository interface {
	FindByID(id string, session *application_specific.Session) (*Plan, *application_specific.ApplicationException)

	Create(plan *Plan, session *application_specific.Session) *application_specific.ApplicationException

	Update(plan *Plan, session *application_specific.Session) *application_specific.ApplicationException
}
