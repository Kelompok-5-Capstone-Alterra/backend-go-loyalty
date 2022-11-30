package routes

import (
	authController "backend-go-loyalty/internal/controller/auth"
	pingController "backend-go-loyalty/internal/controller/ping"
	rewardController "backend-go-loyalty/internal/controller/reward"
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

type rewardRoutes struct {
	rc     rewardController.IRewardController
	router *echo.Echo
}

func NewRewardRoutes(rc rewardController.IRewardController, router *echo.Echo) rewardRoutes {
	return rewardRoutes{
		rc:     rc,
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

func (rrt rewardRoutes) InitEndpoints() {
	reward := rrt.router.Group("/reward")
	reward.GET("", rrt.rc.FindAllReward)
	reward.GET("/:id", rrt.rc.FindRewardById)
	reward.POST("", rrt.rc.CreateReward)
	reward.PUT("/:id", rrt.rc.UpdateReward)
	reward.DELETE("/:id", rrt.rc.DeleteReward)
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
