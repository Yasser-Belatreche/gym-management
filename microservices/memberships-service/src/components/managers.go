package components

import (
	"gym-management-memberships/src/components/memberships"
	"gym-management-memberships/src/lib"
)

func Initialize() {
	lib.InitializeLib()

	memberships.InitializeMembershipsManager(lib.MessagesBroker(), lib.JobsScheduler())
}

func Memberships() memberships.Manager {
	return memberships.NewMembershipsManager(lib.MessagesBroker())
}
