package gin

import (
	g "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gym-management-api-gateway/src/web/gin/v1/controllers/auth"
	"gym-management-api-gateway/src/web/gin/v1/controllers/gyms"
	"gym-management-api-gateway/src/web/gin/v1/controllers/health"
	"gym-management-api-gateway/src/web/gin/v1/controllers/memberships"
	"gym-management-api-gateway/src/web/gin/v1/middlewares"
	"gym-management-api-gateway/src/web/gin/v1/utils"
	"net/http"
	"reflect"
	"strings"
)

func StartWebServer() {
	g.DisableConsoleColor()
	r := g.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			if name == "" {
				name = fld.Tag.Get("form")
			}
			return name
		})
	}

	r.Use(g.CustomRecovery(utils.GlobalErrorHandler))

	r.Use(middlewares.SessionInjectorMiddleware())
	r.Use(middlewares.RequestLoggerMiddleware())

	router := r.Group("/api/v1")

	auth.AuthRouter(router)
	gyms.GymsRouter(router)
	memberships.MembershipsRouter(router)
	health.HealthRoutes(router)

	r.GET("/", func(c *g.Context) {
		c.JSON(http.StatusOK, g.H{"message": "Hello, World!"})
	})

	r.NoRoute(func(c *g.Context) {
		utils.HandleError(c, utils.NewRouteNotFoundError())
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
