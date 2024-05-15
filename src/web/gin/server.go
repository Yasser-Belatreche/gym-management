package gin

import (
	g "github.com/gin-gonic/gin"
	"gym-management/src/web/gin/v1/controllers/auth"
	"gym-management/src/web/gin/v1/controllers/health"
	"gym-management/src/web/gin/v1/middlewares"
	"gym-management/src/web/gin/v1/utils"
	"net/http"
)

func StartWebServer() {
	r := g.Default()

	r.Use(g.CustomRecovery(utils.GlobalErrorHandler))

	r.Use(middlewares.SessionInjectorMiddleware())

	router := r.Group("/api/v1")

	auth.AuthRoutes(router)
	health.HealthRoutes(router)

	r.GET("/", func(c *g.Context) {
		c.JSON(http.StatusOK, g.H{
			"message": "Hello, World!",
		})
	})

	r.NoRoute(func(c *g.Context) {
		utils.HandleError(c, utils.NewRouteNotFoundError())
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
