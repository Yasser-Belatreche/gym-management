package utils

import (
	"github.com/gin-gonic/gin"
	"gym-management-api-gateway/src/components"
	"gym-management-api-gateway/src/components/service_discovery"
	"gym-management-api-gateway/src/lib/primitives/application_specific"
	"net/http"
)

func RedirectToAuthService(c *gin.Context) {
	redirectToService("auth", c)
}

func RedirectToGymsService(c *gin.Context) {
	redirectToService("gyms", c)
}

func RedirectToMembershipsService(c *gin.Context) {
	redirectToService("memberships", c)
}

func redirectToService(name string, c *gin.Context) {
	service, err := getService(name)
	if err != nil {
		HandleError(c, err)
		return
	}

	client := http.Client{}
	req, err := http.NewRequest(c.Request.Method, service.Url+c.Request.RequestURI, c.Request.Body)
	if err != nil {
		HandleError(c, err)
		return
	}

	prepareHeaders(c, service, &req.Header)

	res, err := client.Do(req)
	if err != nil {
		HandleError(c, err)
		return
	}

	CopyResponse(res, c)
}

func getService(name string) (*service_discovery.Service, error) {
	switch name {
	case "auth":
		return components.ServiceDiscovery().GetAuthService()
	case "gyms":
		return components.ServiceDiscovery().GetGymsService()
	case "memberships":
		return components.ServiceDiscovery().GetMembershipsService()
	default:
		return nil, application_specific.NewDeveloperException("CASE_NOT_IMPLEMENTED", "Service "+name+" is not implemented")
	}

}

func prepareHeaders(c *gin.Context, service *service_discovery.Service, headers *http.Header) {
	var session interface{ ToBase64() (string, error) }

	userSession := CheckUserSession(c)
	if userSession {
		session = ExtractUserSession(c)
	} else {
		session = ExtractSession(c)
	}

	base64, err := session.ToBase64()
	if err != nil {
		HandleError(c, err)
		return
	}

	for k, v := range c.Request.Header {
		for _, vv := range v {
			headers.Set(k, vv)
		}
	}

	headers.Set("X-Session", base64)
	headers.Set("X-Api-Secret", service.ApiSecret)
}
