package components

import (
	"gym-management-auth/src/components/auth"
	"gym-management-auth/src/lib"
)

func Initialize() {
	lib.InitializeLib()

	auth.InitializeAuthManager(lib.MessagesBroker())
}

func Auth() auth.Manager {
	return auth.NewAuthManager()
}
