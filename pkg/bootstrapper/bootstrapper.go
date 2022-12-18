package bootstrapper

import (
	authController "backend-go-loyalty/internal/controller/auth"
	categoryController "backend-go-loyalty/internal/controller/category"
	faqController "backend-go-loyalty/internal/controller/faq"
	paymentController "backend-go-loyalty/internal/controller/payment"
	pingController "backend-go-loyalty/internal/controller/ping"
	productController "backend-go-loyalty/internal/controller/product"
	redeemController "backend-go-loyalty/internal/controller/redeem"
	rewardController "backend-go-loyalty/internal/controller/reward"
	transactionController "backend-go-loyalty/internal/controller/transaction"
	userController "backend-go-loyalty/internal/controller/user"
	webhookController "backend-go-loyalty/internal/controller/webhook"
	authRepository "backend-go-loyalty/internal/repository/auth"
	categoryRepository "backend-go-loyalty/internal/repository/category"
	faqRepository "backend-go-loyalty/internal/repository/faq"
	paymentRepository "backend-go-loyalty/internal/repository/payment"
	pointRepository "backend-go-loyalty/internal/repository/point"
	productRepository "backend-go-loyalty/internal/repository/product"
	redeemRepository "backend-go-loyalty/internal/repository/redeem"
	rewardRepository "backend-go-loyalty/internal/repository/reward"
	transactionRepository "backend-go-loyalty/internal/repository/transaction"
	userRepository "backend-go-loyalty/internal/repository/user"
	"backend-go-loyalty/internal/routes"
	authService "backend-go-loyalty/internal/service/auth"
	categoryService "backend-go-loyalty/internal/service/category"
	faqService "backend-go-loyalty/internal/service/faq"
	paymentService "backend-go-loyalty/internal/service/payment"
	pingService "backend-go-loyalty/internal/service/ping"
	productService "backend-go-loyalty/internal/service/product"
	redeemService "backend-go-loyalty/internal/service/redeem"
	rewardService "backend-go-loyalty/internal/service/reward"
	transactionService "backend-go-loyalty/internal/service/transaction"
	userService "backend-go-loyalty/internal/service/user"

	"github.com/labstack/echo/v4"
	"github.com/xendit/xendit-go/client"
	"gorm.io/gorm"
)

func InitEndpoints(router *echo.Echo, db *gorm.DB, xenditAPI *client.API) {
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
	// pointService := pointService.NewPointService(pointRepository)
	// pointController := pointController.NewPointController(pointService)
	// pointRoutes := routes.NewPointRoutes(pointController, router)
	// pointRoutes.InitEndpoints()

	redeemRepository := redeemRepository.NewRedeemRepository(db)
	redeemService := redeemService.NewRedeemService(redeemRepository, pointRepository, rewardRepository, userRepository)
	redeemController := redeemController.NewRedeemController(redeemService)
	redeemRoutes := routes.NewRedeemRoutes(redeemController, router)
	redeemRoutes.InitEndpoints()

	categoryRepository := categoryRepository.NewCategoryRepository(db)
	categoryService := categoryService.NewCategoryService(categoryRepository)
	categoryController := categoryController.NewCategoryController(categoryService)
	categoryRoutes := routes.NewCategoryRoutes(categoryController, router)
	categoryRoutes.InitEndpoints()

	faqRepository := faqRepository.NewFAQRepository(db)
	faqService := faqService.NewFAQService(faqRepository)
	faqController := faqController.NewFAQController(faqService)
	faqRoutes := routes.NewFAQRoutes(faqController, router)
	faqRoutes.InitEndpoints()

	paymentRepository := paymentRepository.NewPaymentRepository(db, xenditAPI.EWallet)

	transactionRepository := transactionRepository.NewTransactionRepository(db)
	transactionService := transactionService.NewTransactionService(transactionRepository, userRepository, productRepository, paymentRepository, pointRepository)
	transactionController := transactionController.NewTransactionController(transactionService)
	transactionRoutes := routes.NewTransactionRoutes(transactionController, router)
	transactionRoutes.InitEndpoints()

	paymentService := paymentService.NewPaymentService(transactionRepository, paymentRepository, userRepository)
	paymentController := paymentController.NewPaymentController(transactionService, paymentService)
	paymentRoutes := routes.NewPaymentRoutes(paymentController, router)
	paymentRoutes.InitEndpoints()

	webhookController := webhookController.NewWebhookController(transactionService)
	webhookRoutes := routes.NewWebhookRoutes(webhookController, router)
	webhookRoutes.InitEndpoints()
}
