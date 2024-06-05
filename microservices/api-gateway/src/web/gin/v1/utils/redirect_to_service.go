package utils

import (
	"github.com/gin-gonic/gin"
	"gym-management-api-gateway/src/components"
	"gym-management-api-gateway/src/lib/primitives/application_specific"
	"io"
	"net/http"
	"os"
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

func redirectToService(service string, c *gin.Context) {
	url, err := getServiceUrl(service)
	if err != nil {
		HandleError(c, err)
		return
	}

	client := http.Client{}
	req, err := http.NewRequest(c.Request.Method, url+c.Request.RequestURI, c.Request.Body)
	if err != nil {
		HandleError(c, err)
		return
	}

	prepareHeaders(c, &req.Header)

	res, err := client.Do(req)
	if err != nil {
		HandleError(c, err)
		return
	}

	copyResponse(res, c)
}

func getServiceUrl(service string) (string, error) {
	switch service {
	case "auth":
		return components.ServiceDiscovery().GetAuthServiceUrl()
	case "gyms":
		return components.ServiceDiscovery().GetGymsServiceUrl()
	case "memberships":
		return components.ServiceDiscovery().GetMembershipsServiceUrl()
	default:
		return "", application_specific.NewDeveloperException("CASE_NOT_IMPLEMENTED", "Service "+service+" is not implemented")
	}

}

func prepareHeaders(c *gin.Context, headers *http.Header) {
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

	apiSecret, exists := os.LookupEnv("API_SECRET")
	if !exists {
		panic("API_SECRET environment variable is not set")
	}

	headers.Set("X-Session", base64)
	headers.Set("X-Api-Secret", apiSecret)
}

func copyResponse(res *http.Response, c *gin.Context) {
	for k, v := range res.Header {
		c.Header(k, v[0])
	}

	c.Status(res.StatusCode)
	io.Copy(c.Writer, res.Body)
}
