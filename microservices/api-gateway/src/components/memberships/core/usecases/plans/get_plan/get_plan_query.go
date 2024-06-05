package get_plan

import "gym-management/src/lib/primitives/application_specific"

type GetPlanQuery struct {
	Id      string
	Session *application_specific.UserSession
}
