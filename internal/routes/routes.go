package routes

import (
	authController "backend-go-loyalty/internal/controller/auth"
	pingController "backend-go-loyalty/internal/controller/ping"
	pointController "backend-go-loyalty/internal/controller/point"
	productController "backend-go-loyalty/internal/controller/product"
	redeemController "backend-go-loyalty/internal/controller/redeem"
	rewardController "backend-go-loyalty/internal/controller/reward"
	userController "backend-go-loyalty/internal/controller/user"
	"backend-go-loyalty/internal/middleware"

	"github.com/labstack/echo/v4"
)

type pingRoutes struct {
	pc     pingController.PingController
	router *echo.Echo
}

type authRoutes struct {
	ac     authController.AuthController
	router *echo.Echo
}

type userRoutes struct {
	uc     userController.UserControllerInterface
	router *echo.Echo
}

type rewardRoutes struct {
	rc     rewardController.IRewardController
	router *echo.Echo
}

type pointRoutes struct {
	pc     pointController.IPointController
	router *echo.Echo
}

func NewPointRoutes(pc pointController.IPointController, router *echo.Echo) pointRoutes {
	return pointRoutes{
		pc:     pc,
		router: router,
	}
}

func NewRewardRoutes(rc rewardController.IRewardController, router *echo.Echo) rewardRoutes {
	return rewardRoutes{
		rc:     rc,
		router: router,
	}
}

type productRoutes struct {
	pc     productController.IProductController
	router *echo.Echo
}

func NewProductRoutes(pc productController.IProductController, router *echo.Echo) productRoutes {
	return productRoutes{
		pc:     pc,
		router: router,
	}
}

func NewPingRoutes(pc pingController.PingController, router *echo.Echo) pingRoutes {
	return pingRoutes{
		pc:     pc,
		router: router,
	}
}

func NewAuthRoutes(ac authController.AuthController, router *echo.Echo) authRoutes {
	return authRoutes{
		ac:     ac,
		router: router,
	}
}

func NewUserRoutes(uc userController.UserControllerInterface, router *echo.Echo) userRoutes {
	return userRoutes{
		uc:     uc,
		router: router,
	}
}

type redeemRoutes struct {
	dc     redeemController.IRedeemController
	router *echo.Echo
}

func NewRedeemRoutes(dc redeemController.IRedeemController, router *echo.Echo) redeemRoutes {
	return redeemRoutes{
		dc:     dc,
		router: router,
	}
}

func (prt pointRoutes) InitEndpoints() {
	point := prt.router.Group("/points")
	adminPoints := prt.router.Group("/admin/points", middleware.ValidateAdminJWT)

	adminPoints.GET("", prt.pc.HandleGetAllPoint)
	point.GET("/:id", prt.pc.HandleGetPointByID)
}

func (rrt rewardRoutes) InitEndpoints() {
	reward := rrt.router.Group("/rewards")
	reward.GET("", rrt.rc.FindAllReward)
	reward.GET("/:id", rrt.rc.FindRewardById)
	reward.POST("", rrt.rc.CreateReward)

	adminReward := rrt.router.Group("/admin/rewards", middleware.ValidateAdminJWT)
	adminReward.PUT("/:id", rrt.rc.UpdateReward)
	adminReward.DELETE("/:id", rrt.rc.DeleteReward)
}

func (prt pingRoutes) InitEndpoints() {
	ping := prt.router.Group("/ping")
	ping.GET("", prt.pc.HandlePing)
}

func (art authRoutes) InitEndpoints() {
	auth := art.router.Group("/auth")
	auth.POST("/signin", art.ac.HandleLogin)
	auth.POST("/signup", art.ac.HandleSignUp)

	token := auth.Group("/token")
	token.POST("/refresh", art.ac.HandleRefreshToken)

	otp := auth.Group("/otp")
	otp.POST("/otp/validate", art.ac.HandleValidateOTP)
	otp.POST("/otp/resend", art.ac.HandleRequestNewOTP)
}

func (urt userRoutes) InitEndpoints() {
	user := urt.router.Group("/users", middleware.ValidateJWT)
	user.PUT("/change-password", urt.uc.HandleChangePassword)
	user.PUT("", urt.uc.HandleUpdateData)

	admin := urt.router.Group("/users", middleware.ValidateAdminJWT)
	admin.GET("", urt.uc.HandleGetAllUser)
	admin.GET("/:id", urt.uc.HandleGetUserByID)
	admin.PUT("/:id", urt.uc.HandleUpdateCustomerData)
	admin.DELETE("/:id", urt.uc.HandleDeleteCustomerData)
}

func (prt productRoutes) InitEndpoints() {
	product := prt.router.Group("/products")
	product.GET("", prt.pc.GetAll)
	product.GET("/:id", prt.pc.GetProductById)
	product.POST("", prt.pc.InsertProduct)
	product.PUT("/:id", prt.pc.UpdateProduct)
	product.DELETE("/:id", prt.pc.DeleteProduct)
}

func (drt redeemRoutes) InitEndpoints() {
	redeem := drt.router.Group("/redeems")
	redeem.GET("", drt.dc.GetAllRedeem)
	redeem.GET("/:id", drt.dc.GetRedeemByID)
	redeem.POST("", drt.dc.CreateRedeem)
	redeem.PUT("/:id", drt.dc.UpdateRedeem)
	redeem.DELETE("/:id", drt.dc.DeleteRedeem)
}
