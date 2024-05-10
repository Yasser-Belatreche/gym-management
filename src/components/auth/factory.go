package auth

import (
	"gym-management/src/lib/messages_broker"
	"os"
)

var manager Manager = nil

func NewAuthManager() Manager {
	secret, found := os.LookupEnv("JWT_SECRET")
	if !found {
		panic("JWT_SECRET env var is required")
	}

	if manager == nil {
		manager = &Facade{
			UserRepository: nil,
			JwtSecret:      secret,
		}
	}

	return manager
}

func InitializeAuthManager(broker messages_broker.MessagesBroker) {
	secret, found := os.LookupEnv("JWT_SECRET")
	if !found {
		panic("JWT_SECRET env var is required")
	}

	initialize(broker, &Facade{
		UserRepository: nil,
		JwtSecret:      secret,
	})
}
