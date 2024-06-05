package health

import (
	"github.com/gin-gonic/gin"
	"gym-management-memberships/src/lib"
)

func HealthRoutes(router *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) {
		broker := lib.MessagesBroker().HealthCheck()
		persistence := lib.Persistence().HealthCheck()

		status := "UP"

		if broker.Status != "UP" {
			status = "DOWN"
		}

		if persistence.Status != "UP" {
			status = "DOWN"
		}

		c.JSON(200, gin.H{
			"status":         status,
			"persistence":    persistence,
			"messagesBroker": broker,
		})
	})
}
