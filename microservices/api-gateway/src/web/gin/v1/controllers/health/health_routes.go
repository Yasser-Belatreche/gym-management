package health

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gym-management-api-gateway/src/components"
	"gym-management-api-gateway/src/web/gin/v1/utils"
	"net/http"
	"os"
)

func HealthRoutes(router *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) {
		auth, err := getAuthHealth(c)
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		memberships, err := getMembershipsHealth(c)
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		gyms, err := getGymsHealth(c)
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		serviceDiscovery, err := components.ServiceDiscovery().GetHealth()
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		status := "UP"

		if auth["status"] != "UP" {
			status = "DOWN"
		}
		if gyms["status"] != "UP" {
			status = "DOWN"
		}
		if memberships["status"] != "UP" {
			status = "DOWN"
		}
		if serviceDiscovery["status"] != "UP" {
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
	authUrl, err := components.ServiceDiscovery().GetAuthServiceUrl()
	if err != nil {
		return nil, err
	}

	url := authUrl + "/api/v1/health"

	return getHealth(url, c)
}

func getGymsHealth(c *gin.Context) (map[string]interface{}, error) {
	authUrl, err := components.ServiceDiscovery().GetGymsServiceUrl()
	if err != nil {
		return nil, err
	}

	url := authUrl + "/api/v1/health"

	return getHealth(url, c)
}

func getMembershipsHealth(c *gin.Context) (map[string]interface{}, error) {
	authUrl, err := components.ServiceDiscovery().GetMembershipsServiceUrl()
	if err != nil {
		return nil, err
	}

	url := authUrl + "/api/v1/health"

	return getHealth(url, c)
}

func getHealth(url string, c *gin.Context) (map[string]interface{}, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
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

	apiSecret, exists := os.LookupEnv("API_SECRET")
	if !exists {
		panic("API_SECRET environment variable is not set")
	}

	req.Header.Set("X-Session", base64)
	req.Header.Set("X-Api-Secret", apiSecret)

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
