package gyms

var manager Manager = nil

func NewGymsManager() Manager {
	if manager == nil {
		manager = &AuthorizationDecorator{
			manager: &Facade{
				EmailService:       nil,
				EventsPublisher:    nil,
				GymOwnerRepository: nil,
			},
		}
	}

	return manager
}
