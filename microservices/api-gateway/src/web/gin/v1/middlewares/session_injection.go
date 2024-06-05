package middlewares

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/lib/primitives/application_specific"
	"gym-management/src/lib/primitives/generic"
)

func SessionInjectorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		correlationId := c.GetHeader("X-Correlation-Id")
		if correlationId == "" {
			correlationId = generic.GenerateUUID()
		}

		c.Set("session", application_specific.NewSessionWithCorrelationId(correlationId))

		c.Next()
	}
}
