package components

import (
	"gym-management/src/components/auth"
	"gym-management/src/components/gyms"
	"gym-management/src/lib"
)

func InitManagers() {
	auth.InitializeAuthManager(lib.MessagesBroker())
}

func Gyms() gyms.Manager {
	return gyms.NewGymsManager(lib.MessagesBroker())
}

func Auth() auth.Manager {
	return auth.NewAuthManager()
}
