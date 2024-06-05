package auth

import (
	"gym-management-auth/src/components/auth/core/domain"
	"gym-management-auth/src/components/auth/infra"
	"gym-management-auth/src/lib/messages_broker"
	"os"
)

var manager Manager = nil
var facade *Facade = nil

func getFacade(userRepository domain.UserRepository, jwtSecret string) *Facade {
	if facade != nil {
		return facade
	}

	facade = &Facade{
		UserRepository: userRepository,
		JwtSecret:      jwtSecret,
	}

	return facade
}

func NewAuthManager() Manager {
	secret, found := os.LookupEnv("JWT_SECRET")
	if !found {
		panic("JWT_SECRET env var is required")
	}

	if manager == nil {
		facade = getFacade(&infra.GormUserRepository{}, secret)
		manager = facade
	}

	return manager
}

func InitializeAuthManager(broker messages_broker.MessagesBroker) {
	secret, found := os.LookupEnv("JWT_SECRET")
	if !found {
		panic("JWT_SECRET env var is required")
	}

	initialize(broker, getFacade(&infra.GormUserRepository{}, secret))
}
