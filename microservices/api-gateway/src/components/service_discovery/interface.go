package service_discovery

type ServiceDiscovery interface {
	GetAuthService() (*Service, error)

	GetMembershipsService() (*Service, error)

	GetGymsService() (*Service, error)

	GetHealth() (map[string]interface{}, error)
}

type Service struct {
	Url       string
	ApiSecret string
}
