package components

import (
	"gym-management/src/components/gyms"
	"gym-management/src/lib"
)

func InitManagers() {}

func Gyms() gyms.Manager {
	return gyms.NewGymsManager(lib.MessagesBroker())
}
