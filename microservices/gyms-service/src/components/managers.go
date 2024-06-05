package components

import (
	"gym-management-gyms/src/components/gyms"
	"gym-management-gyms/src/lib"
)

func Initialize() {
	lib.InitializeLib()

	gyms.InitializeGymsManager()
}

func Gyms() gyms.Manager {
	return gyms.NewGymsManager(lib.MessagesBroker())
}
