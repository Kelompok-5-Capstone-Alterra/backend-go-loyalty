package routes

import (
	authController "backend-go-loyalty/internal/controller/auth"
	pingController "backend-go-loyalty/internal/controller/ping"
	productController "backend-go-loyalty/internal/controller/product"
	userController "backend-go-loyalty/internal/controller/user"

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

func (prt pingRoutes) InitEndpoints() {
	ping := prt.router.Group("/ping")
	ping.GET("", prt.pc.HandlePing)
}

func (art authRoutes) InitEndpoints() {
	auth := art.router.Group("/auth")
	auth.POST("/signin", art.ac.HandleLogin)
	auth.POST("/signup", art.ac.HandleSignUp)
	auth.POST("/otp/validate", art.ac.HandleValidateOTP)
	auth.POST("/token/refresh", art.ac.HandleRefreshToken)
	auth.POST("/otp/resend", art.ac.HandleRequestNewOTP)
}

func (urt userRoutes) InitEndpoints() {
	user := urt.router.Group("/user")
	user.PUT("/change-password", urt.uc.HandleChangePassword)
	user.PUT("", urt.uc.HandleUpdateData)
}

func (prt productRoutes) InitEndpoints() {
	product := prt.router.Group("/product")
	product.GET("", prt.pc.GetAll)
	product.GET("/:id", prt.pc.GetProductById)
	product.POST("", prt.pc.InsertProduct)
	product.PUT("/:id", prt.pc.UpdateProduct)
	product.DELETE("/:id", prt.pc.DeleteProduct)
}
