package health

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/lib"
)

func HealthRoutes(router *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) {
		health := lib.Persistence().HealthCheck()

		status := "UP"

		if health.Status != "UP" {
			status = "DOWN"
		}

		c.JSON(200, gin.H{
			"status":      status,
			"persistence": health,
		})
	})
}
