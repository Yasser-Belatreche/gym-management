package utils

import (
	"github.com/gin-gonic/gin"
	"gym-management-auth/src/lib/primitives/application_specific"
)

func ExtractSession(c *gin.Context) *application_specific.Session {
	var session *application_specific.Session

	val, _ := c.Get("session")

	switch s := val.(type) {
	case *application_specific.Session:
		session = s
	case *application_specific.UserSession:
		session = s.Session
	default:
		panic("Should register the SessionExtractorMiddleware first")
	}

	return session
}

func ExtractUserSession(c *gin.Context) *application_specific.UserSession {
	var session *application_specific.UserSession

	val, exists := c.Get("session")
	if !exists {
		panic("Should register the SessionExtractorMiddleware first")
	}

	switch s := val.(type) {
	case *application_specific.Session:
		panic("Should check if the session is a user session first")
	case *application_specific.UserSession:
		session = s
	default:
		panic("Should register the SessionExtractorMiddleware first")
	}

	return session
}

func CheckUserSession(c *gin.Context) bool {
	session, exists := c.Get("session")
	if !exists {
		return false
	}

	_, ok := session.(*application_specific.UserSession)
	return ok
}

func CheckSession(c *gin.Context) bool {
	_, exists := c.Get("session")
	return exists
}
