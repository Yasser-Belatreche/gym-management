package delete_plan

import "gym-management/src/lib/primitives/application_specific"

type DeletePlanCommand struct {
	Id      string
	Session *application_specific.UserSession
}
