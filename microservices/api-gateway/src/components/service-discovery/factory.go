package service_discovery

import "os"

var instance ServiceDiscovery

func NewServiceDiscovery() ServiceDiscovery {
	serviceDiscoveryUrl, exists := os.LookupEnv("SERVICE_DISCOVERY_URL")
	if !exists {
		panic("SERVICE_DISCOVERY_URL env var is required")
	}

	apiSecret, exists := os.LookupEnv("API_SECRET")
	if !exists {
		panic("API_SECRET env var is required")
	}

	if instance == nil {
		instance = &facade{
			serviceDiscoveryUrl: serviceDiscoveryUrl,
			apiSecret:           apiSecret,
		}
	}

	return instance
}
