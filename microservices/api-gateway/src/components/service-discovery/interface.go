package service_discovery

type ServiceDiscovery interface {
	GetAuthServiceUrl() (string, error)

	GetMembershipsServiceUrl() (string, error)

	GetGymsServiceUrl() (string, error)

	GetHealth() (map[string]interface{}, error)
}
