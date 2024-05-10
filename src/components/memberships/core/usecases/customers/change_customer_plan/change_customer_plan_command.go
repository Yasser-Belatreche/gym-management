package change_customer_plan

import (
	"gym-management/src/lib/primitives/application_specific"
	"time"
)

type ChangeCustomerPlanCommand struct {
	CustomerId string
	PlanId     string
	EndDate    *time.Time
	Session    *application_specific.UserSession
}
