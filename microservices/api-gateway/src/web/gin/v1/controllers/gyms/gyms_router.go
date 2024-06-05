package gyms

import (
	"github.com/gin-gonic/gin"
	"gym-management-api-gateway/src/web/gin/v1/middlewares"
	"gym-management-api-gateway/src/web/gin/v1/utils"
)

func GymsRouter(router *gin.RouterGroup) {
	g := router.Group("/owners")

	g.Use(middlewares.AuthMiddleware())
	{
		g.GET("/:ownerId", utils.RedirectToGymsService)
		g.GET("/", utils.RedirectToGymsService)
		g.POST("/", utils.RedirectToGymsService)
		g.PUT("/:ownerId", utils.RedirectToGymsService)
		g.DELETE("/:ownerId", utils.RedirectToGymsService)
		g.PATCH("/:ownerId/restrict", utils.RedirectToGymsService)
		g.PATCH("/:ownerId/unrestrict", utils.RedirectToGymsService)
	}

	gyms := g.Group("/:ownerId/gyms")
	{
		gyms.GET("/:gymId", utils.RedirectToGymsService)
		gyms.GET("/", utils.RedirectToGymsService)
		gyms.POST("/", utils.RedirectToGymsService)
		gyms.PUT("/:gymId", utils.RedirectToGymsService)
		gyms.DELETE("/:gymId", utils.RedirectToGymsService)
		gyms.PATCH("/:gymId/enable", utils.RedirectToGymsService)
		gyms.PATCH("/:gymId/disable", utils.RedirectToGymsService)
	}
}
