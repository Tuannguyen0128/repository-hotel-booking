package routes

import (
	"github.com/gin-gonic/gin"
	"repository-hotel-booking/internal/app/service"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(c *gin.Context)
}

//	func SetupRoutes(c *gin.Engine) {
//		for _, router := range Load() {
//			c.Handle(router.Method,router.Uri, router.Handler)
//		}
//	}
func Load(s *service.Service) []Route {
	accountRoutes := AccountRoutes(s)
	routes := append(accountRoutes)
	//routes = append(routes, loginRoutes...)
	return routes
}
func SetupRoutes(g *gin.Engine, s *service.Service) {
	for _, router := range Load(s) {
		g.Handle(router.Method, "/api"+router.Uri,
			router.Handler,
		)
	}
}
