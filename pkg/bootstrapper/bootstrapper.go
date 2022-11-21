package bootstrapper

import (
	pingController "backend-go-loyalty/internal/controller/ping"
	"backend-go-loyalty/internal/routes"
	pingService "backend-go-loyalty/internal/service/ping"

	"github.com/labstack/echo/v4"
)

func InitEndpoints(router *echo.Echo) {
	pingService := pingService.NewPingService()
	pingController := pingController.NewPingController(pingService)
	pingRoutes := routes.NewPingRoutes(pingController, router)
	pingRoutes.InitEndpoints()

}
