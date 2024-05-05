package domain

import "gym-management/src/lib/primitives/application_specific"

type PlanRepository interface {
	FindByID(id string) (*Plan, *application_specific.ApplicationException)

	Create(plan *Plan) *application_specific.ApplicationException

	Update(plan *Plan) *application_specific.ApplicationException
}
