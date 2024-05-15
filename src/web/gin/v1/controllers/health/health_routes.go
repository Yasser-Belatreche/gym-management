package health

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/lib"
)

func HealthRoutes(router *gin.RouterGroup) {
	router.GET("/health", func(c *gin.Context) {
		health := lib.Persistence().HealthCheck()

		c.JSON(200, gin.H{
			"persistence": health,
		})
	})
}
