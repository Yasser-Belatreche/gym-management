package utils

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/lib/primitives/application_specific"
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
		panic("Should register the SessionInjectorMiddleware first")
	}

	return session
}

func ExtractUserSession(c *gin.Context) *application_specific.UserSession {
	var session *application_specific.UserSession

	val, _ := c.Get("session")

	switch s := val.(type) {
	case *application_specific.Session:
		panic("Should register the AuthMiddleware after the SessionInjectorMiddleware, and the end point should not be public")
	case *application_specific.UserSession:
		session = s
	default:
		panic("Should register the SessionInjectorMiddleware first")
	}

	return session
}
