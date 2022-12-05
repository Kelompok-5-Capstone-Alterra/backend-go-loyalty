package bootstrapper

import (
	authController "backend-go-loyalty/internal/controller/auth"
	pingController "backend-go-loyalty/internal/controller/ping"
	pointController "backend-go-loyalty/internal/controller/point"
	productController "backend-go-loyalty/internal/controller/product"
	redeemController "backend-go-loyalty/internal/controller/redeem"
	rewardController "backend-go-loyalty/internal/controller/reward"
	userController "backend-go-loyalty/internal/controller/user"
	authRepository "backend-go-loyalty/internal/repository/auth"
	pointRepository "backend-go-loyalty/internal/repository/point"
	productRepository "backend-go-loyalty/internal/repository/product"
	redeemRepository "backend-go-loyalty/internal/repository/redeem"
	rewardRepository "backend-go-loyalty/internal/repository/reward"
	userRepository "backend-go-loyalty/internal/repository/user"
	"backend-go-loyalty/internal/routes"
	authService "backend-go-loyalty/internal/service/auth"
	pingService "backend-go-loyalty/internal/service/ping"
	pointService "backend-go-loyalty/internal/service/point"
	productService "backend-go-loyalty/internal/service/product"
	redeemService "backend-go-loyalty/internal/service/redeem"
	rewardService "backend-go-loyalty/internal/service/reward"
	userService "backend-go-loyalty/internal/service/user"

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

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userController := userController.NewUserController(userService)
	userRoutes := routes.NewUserRoutes(userController, router)
	userRoutes.InitEndpoints()

	rewardRepository := rewardRepository.NewRewardRepository(db)
	rewardService := rewardService.NewRewardService(rewardRepository)
	rewardController := rewardController.NewRewardController(rewardService)
	rewardRoutes := routes.NewRewardRoutes(rewardController, router)
	rewardRoutes.InitEndpoints()

	productRepository := productRepository.NewProductRepository(db)
	productService := productService.NewProductService(productRepository)
	productController := productController.NewProductController(productService)
	productRoutes := routes.NewProductRoutes(productController, router)
	productRoutes.InitEndpoints()

	pointRepository := pointRepository.NewPointRepository(db)
	pointService := pointService.NewPointService(pointRepository)
	pointController := pointController.NewPointController(pointService)
	pointRoutes := routes.NewPointRoutes(pointController, router)
	pointRoutes.InitEndpoints()

	redeemRepository := redeemRepository.NewRedeemRepository(db)
	redeemService := redeemService.NewRedeemService(redeemRepository, pointRepository, rewardRepository)
	redeemController := redeemController.NewRedeemController(redeemService)
	redeemRoutes := routes.NewRedeemRoutes(redeemController, router)
	redeemRoutes.InitEndpoints()

}
