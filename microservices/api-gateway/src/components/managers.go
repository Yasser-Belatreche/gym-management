package components

import (
	service_discovery "gym-management-api-gateway/src/components/service_discovery"
)

func Initialize() {
}

func ServiceDiscovery() service_discovery.ServiceDiscovery {
	return service_discovery.NewServiceDiscovery()
}
