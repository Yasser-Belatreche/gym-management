package components

import "gym-management/src/components/gyms"

func InitManagers() {}

func Gyms() gyms.Manager {
	return gyms.NewGymsManager()
}
