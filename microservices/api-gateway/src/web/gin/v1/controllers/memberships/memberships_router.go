package memberships

import (
	"github.com/gin-gonic/gin"
	"gym-management-api-gateway/src/web/gin/v1/middlewares"
	"gym-management-api-gateway/src/web/gin/v1/utils"
)

func MembershipsRouter(r *gin.RouterGroup) {
	g := r.Group("/owners/:ownerId/gyms/:gymId")

	g.Use(middlewares.AuthMiddleware())

	plans := g.Group("/plans")
	{
		plans.POST("/", utils.RedirectToMembershipsService)
		plans.PUT("/:planId", utils.RedirectToMembershipsService)
		plans.DELETE("/:planId", utils.RedirectToMembershipsService)
		plans.GET("/", utils.RedirectToMembershipsService)
		plans.GET("/:planId", utils.RedirectToMembershipsService)
	}

	memberships := g.Group("/memberships")
	{
		memberships.PATCH("/:membershipId/cancel", utils.RedirectToMembershipsService)
		memberships.PATCH("/:membershipId/renew", utils.RedirectToMembershipsService)
		memberships.GET("/", utils.RedirectToMembershipsService)
		memberships.GET("/:membershipId", utils.RedirectToMembershipsService)
		memberships.GET("/:membershipId/badge", utils.RedirectToMembershipsService)
	}

	bills := g.Group("/memberships/:membershipId/bills")
	{
		bills.PATCH("/:billId/paid", utils.RedirectToMembershipsService)
		bills.GET("/", utils.RedirectToMembershipsService)
		bills.GET("/:billId", utils.RedirectToMembershipsService)
	}

	trainingSessions := g.Group("/memberships/:membershipId/training-sessions")
	{
		trainingSessions.POST("/", utils.RedirectToMembershipsService)
		trainingSessions.PATCH("/:sessionId/end", utils.RedirectToMembershipsService)
		trainingSessions.GET("/", utils.RedirectToMembershipsService)
		trainingSessions.GET("/:sessionId", utils.RedirectToMembershipsService)
	}

	customers := g.Group("/customers")
	{
		customers.POST("/", utils.RedirectToMembershipsService)
		customers.PUT("/:customerId", utils.RedirectToMembershipsService)
		customers.PATCH("/:customerId/plan", utils.RedirectToMembershipsService)
		customers.PATCH("/:customerId/restrict", utils.RedirectToMembershipsService)
		customers.PATCH("/:customerId/unrestrict", utils.RedirectToMembershipsService)
		customers.DELETE("/:customerId", utils.RedirectToMembershipsService)
		customers.GET("/", utils.RedirectToMembershipsService)
		customers.GET("/:customerId", utils.RedirectToMembershipsService)
	}
}
