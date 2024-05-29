package memberships

import (
	"github.com/gin-gonic/gin"
	"gym-management/src/web/gin/v1/middlewares"
)

func MembershipsRouter(r *gin.RouterGroup) {
	g := r.Group("/owners/:ownerId/gyms/:gymId")

	g.Use(middlewares.AuthMiddleware())

	plans := g.Group("/plans")
	{
		plans.POST("/", middlewares.TransactionMiddleware(), CreatePlanHandler)
		plans.PUT("/:planId", middlewares.TransactionMiddleware(), UpdatePlanHandler)
		plans.DELETE("/:planId", middlewares.TransactionMiddleware(), DeletePlanHandler)
		plans.GET("/", GetPlansHandler)
		plans.GET("/:planId", GetPlanHandler)
	}

	memberships := g.Group("/memberships")
	{
		memberships.PATCH("/:membershipId/cancel", middlewares.TransactionMiddleware(), CancelMembershipHandler)
		memberships.PATCH("/:membershipId/renew", middlewares.TransactionMiddleware(), RenewMembershipHandler)
		memberships.GET("/", GetMembershipsHandler)
		memberships.GET("/:membershipId", GetMembershipHandler)
		memberships.GET("/:membershipId/badge", GetMembershipHandler)
	}

	bills := g.Group("/memberships/:membershipId/bills")
	{
		bills.PATCH("/:billId/paid", middlewares.TransactionMiddleware(), MarkBillAsPaid)
		bills.GET("/", GetBillsHandler)
		bills.GET("/:billId", GetBillHandler)
	}

	trainingSessions := g.Group("/memberships/:membershipId/training-sessions")
	{
		trainingSessions.POST("/", middlewares.TransactionMiddleware(), StartTrainingSessionHandler)
		trainingSessions.PATCH("/:sessionId/end", middlewares.TransactionMiddleware(), EndTrainingSessionHandler)
		trainingSessions.GET("/", GetTrainingSessionsHandler)
		trainingSessions.GET("/:sessionId", GetTrainingSessionHandler)
	}

	customers := g.Group("/customers")
	{
		customers.POST("/", middlewares.TransactionMiddleware(), CreateCustomerHandler)
		customers.PUT("/:customerId", middlewares.TransactionMiddleware(), UpdateCustomerHandler)
		customers.PATCH("/:customerId/plan", middlewares.TransactionMiddleware(), ChangeCustomerPlanHandler)
		customers.PATCH("/:customerId/restrict", middlewares.TransactionMiddleware(), RestrictCustomerHandler)
		customers.PATCH("/:customerId/unrestrict", middlewares.TransactionMiddleware(), UnrestrictCustomerHandler)
		customers.DELETE("/:customerId", middlewares.TransactionMiddleware(), DeleteCustomerHandler)
		customers.GET("/", GetCustomersHandler)
		customers.GET("/:customerId", GetCustomerHandler)
	}
}
