package health

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gym-management-api-gateway/src/components"
	"gym-management-api-gateway/src/components/service_discovery"
	"gym-management-api-gateway/src/web/gin/v1/utils"
	"net/http"
)

func HealthRoutes(router *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) {
		serviceDiscovery, err := components.ServiceDiscovery().GetHealth()
		if err != nil {
			serviceDiscovery = map[string]interface{}{
				"status":  "DOWN",
				"message": err.Error(),
			}
		}

		auth, err := getAuthHealth(c)
		if err != nil {
			auth = map[string]interface{}{
				"status":  "DOWN",
				"message": err.Error(),
			}
		}

		memberships, err := getMembershipsHealth(c)
		if err != nil {
			memberships = map[string]interface{}{
				"status":  "DOWN",
				"message": err.Error(),
			}
		}

		gyms, err := getGymsHealth(c)
		if err != nil {
			gyms = map[string]interface{}{
				"status":  "DOWN",
				"message": err.Error(),
			}
		}

		status := "UP"

		if serviceDiscovery["status"] != "UP" {
			status = "DOWN"
		}
		if auth["status"] != "UP" {
			status = "DOWN"
		}
		if gyms["status"] != "UP" {
			status = "DOWN"
		}
		if memberships["status"] != "UP" {
			status = "DOWN"
		}

		c.JSON(200, gin.H{
			"status":              status,
			"auth-service":        auth,
			"gyms-service":        gyms,
			"memberships-service": memberships,
			"service-discovery":   serviceDiscovery,
		})
	})
}

func getAuthHealth(c *gin.Context) (map[string]interface{}, error) {
	service, err := components.ServiceDiscovery().GetAuthService()
	if err != nil {
		return nil, err
	}

	service.Url = service.Url + "/api/v1/health"

	return getHealth(service, c)
}

func getGymsHealth(c *gin.Context) (map[string]interface{}, error) {
	service, err := components.ServiceDiscovery().GetGymsService()
	if err != nil {
		return nil, err
	}

	service.Url = service.Url + "/api/v1/health"

	return getHealth(service, c)
}

func getMembershipsHealth(c *gin.Context) (map[string]interface{}, error) {
	service, err := components.ServiceDiscovery().GetMembershipsService()
	if err != nil {
		return nil, err
	}

	service.Url = service.Url + "/api/v1/health"

	return getHealth(service, c)
}

func getHealth(service *service_discovery.Service, c *gin.Context) (map[string]interface{}, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", service.Url, nil)
	if err != nil {
		return nil, err
	}

	var session interface{ ToBase64() (string, error) }

	userSession := utils.CheckUserSession(c)
	if userSession {
		session = utils.ExtractUserSession(c)
	} else {
		session = utils.ExtractSession(c)
	}

	base64, err := session.ToBase64()
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Session", base64)
	req.Header.Set("X-Api-Secret", service.ApiSecret)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
