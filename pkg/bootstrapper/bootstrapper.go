package bootstrapper

import (
	authController "backend-go-loyalty/internal/controller/auth"
	pingController "backend-go-loyalty/internal/controller/ping"
	authRepository "backend-go-loyalty/internal/repository/auth"
	"backend-go-loyalty/internal/routes"
	authService "backend-go-loyalty/internal/service/auth"
	pingService "backend-go-loyalty/internal/service/ping"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitEndpoints(router *echo.Echo, db *gorm.DB) {
	pingService := pingService.NewPingService()
	pingController := pingController.NewPingController(pingService)
	pingRoutes := routes.NewPingRoutes(pingController, router)
	pingRoutes.InitEndpoints()

	authRepository := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepository)
	authController := authController.NewAuthController(authService)
	authRoutes := routes.NewAuthRoutes(authController, router)
	authRoutes.InitEndpoints()
}
