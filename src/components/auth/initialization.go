package auth

import (
	"gym-management/src/components/auth/core/event_handlers/customers"
	"gym-management/src/components/auth/core/event_handlers/gym_owners"
	"gym-management/src/components/auth/core/registered_answers"
	"gym-management/src/components/auth/core/usecases/create_admin"
	"gym-management/src/lib/messages_broker"
	"os"
)

func Init(broker messages_broker.MessagesBroker, facade *Facade) {
	createDefaultAdmin(facade)

	broker.RegisterAnswer(registered_answers.BuildIsEmailUsedAnswer(facade.UserRepository))

	broker.Subscribe(gym_owners.BuildGymOwnerCreatedEventHandler(facade.UserRepository))
	broker.Subscribe(gym_owners.BuildGymOwnerUpdatedEventHandler(facade.UserRepository))
	broker.Subscribe(gym_owners.BuildGymOwnerDeletedEventHandler(facade.UserRepository))

	broker.Subscribe(customers.BuildCustomerCreatedEventHandler(facade.UserRepository))
	broker.Subscribe(customers.BuildCustomerUpdatedEventHandler(facade.UserRepository))
	broker.Subscribe(customers.BuildCustomerDeletedEventHandler(facade.UserRepository))
}

func createDefaultAdmin(facade *Facade) {
	email, found := os.LookupEnv("ADMIN_EMAIL")
	if !found {
		panic("ADMIN_EMAIL env var is required")
	}

	password, found := os.LookupEnv("ADMIN_PASSWORD")
	if !found {
		panic("ADMIN_PASSWORD env var is required")
	}

	_, err := facade.CreateAdmin(&create_admin.CreateAdminCommand{
		FirstName: "Yasser",
		LastName:  "Belatreche",
		Email:     email,
		Password:  password,
		Phone:     "07 98 98 09 75",
	})

	if err == nil {
	}
}
