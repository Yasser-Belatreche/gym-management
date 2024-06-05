package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gym-management-api-gateway/src/components"
	"gym-management-api-gateway/src/lib/primitives/application_specific"
	"gym-management-api-gateway/src/web/gin/v1/utils"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := utils.ExtractSession(c)
		sessionBase64, err := session.ToBase64()
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		authHeader := c.GetHeader("Authorization")

		httpClient := http.Client{}

		service, err := components.ServiceDiscovery().GetAuthService()
		request, err := http.NewRequest("GET", service.Url+"/api/v1/auth/session", nil)
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		request.Header.Add("Authorization", authHeader)
		request.Header.Add("X-Session", sessionBase64)
		request.Header.Add("X-Api-Secret", service.ApiSecret)

		resp, err := httpClient.Do(request)
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			utils.CopyResponse(resp, c)
			c.Abort()
			return
		}

		var userSession application_specific.UserSession

		err = json.NewDecoder(resp.Body).Decode(&userSession)
		if err != nil {
			utils.HandleError(c, err)
			return
		}

		c.Set("session", &userSession)

		c.Next()
	}
}
