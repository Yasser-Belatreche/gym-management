package components

import (
	"gym-management/src/components/auth"
	"gym-management/src/components/gyms"
	"gym-management/src/components/memberships"
	"gym-management/src/lib"
)

func InitManagers() {
	auth.InitializeAuthManager(lib.MessagesBroker())
	gyms.InitializeGymsManager()
	memberships.InitializeMembershipsManager(lib.MessagesBroker(), lib.JobsScheduler())
}

func Auth() auth.Manager {
	return auth.NewAuthManager()
}

func Gyms() gyms.Manager {
	return gyms.NewGymsManager(lib.MessagesBroker())
}

func Memberships() memberships.Manager {
	return memberships.NewMembershipsManager(lib.MessagesBroker())
}
