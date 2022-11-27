package routes

import (
	authController "backend-go-loyalty/internal/controller/auth"
	pingController "backend-go-loyalty/internal/controller/ping"
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
	user := urt.router.Group("/user", middleware.ValidateJWT)
	user.PUT("/change-password", urt.uc.HandleChangePassword)
	user.PUT("", urt.uc.HandleUpdateData)
}
