package routes

import (
	pingController "backend-go-loyalty/internal/controller/ping"

	"github.com/labstack/echo/v4"
)

type pingRoutes struct {
	pc     pingController.PingController
	router *echo.Echo
}

func NewPingRoutes(pc pingController.PingController, router *echo.Echo) pingRoutes {
	return pingRoutes{
		pc:     pc,
		router: router,
	}
}

func (prt pingRoutes) InitEndpoints() {
	ping := prt.router.Group("/ping")
	ping.GET("", prt.pc.HandlePing)
}
