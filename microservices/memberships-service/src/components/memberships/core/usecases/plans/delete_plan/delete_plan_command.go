package delete_plan

import "gym-management-memberships/src/lib/primitives/application_specific"

type DeletePlanCommand struct {
	Id      string
	Session *application_specific.UserSession
}
