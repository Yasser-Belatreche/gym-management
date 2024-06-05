package components

import (
	service_discovery "gym-management-api-gateway/src/components/service-discovery"
)

func Initialize() {
}

func ServiceDiscovery() service_discovery.ServiceDiscovery {
	return service_discovery.NewServiceDiscovery()
}
