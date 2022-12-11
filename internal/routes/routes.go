package routes

import (
	authController "backend-go-loyalty/internal/controller/auth"
	categoryController "backend-go-loyalty/internal/controller/category"
	faqController "backend-go-loyalty/internal/controller/faq"
	pingController "backend-go-loyalty/internal/controller/ping"
	pointController "backend-go-loyalty/internal/controller/point"
	productController "backend-go-loyalty/internal/controller/product"
	redeemController "backend-go-loyalty/internal/controller/redeem"
	rewardController "backend-go-loyalty/internal/controller/reward"
	userController "backend-go-loyalty/internal/controller/user"
	"backend-go-loyalty/internal/middleware"

	"github.com/labstack/echo/v4"
)

type faqRoutes struct {
	fc     faqController.IFaqController
	router *echo.Echo
}

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

type categoryRoutes struct {
	cc     categoryController.ICategoryController
	router *echo.Echo
}

func NewFAQRoutes(fc faqController.IFaqController, router *echo.Echo) faqRoutes {
	return faqRoutes{
		fc:     fc,
		router: router,
	}
}

func NewCategoryRoutes(cc categoryController.ICategoryController, router *echo.Echo) categoryRoutes {
	return categoryRoutes{
		cc:     cc,
		router: router,
	}
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

func (frt faqRoutes) InitEndpoints() {
	faq := frt.router.Group("/faqs")
	faq.GET("", frt.fc.HandleGetAllFAQByKeyword)
	faq.GET("/:id", frt.fc.HandleGetFAQByID)

	adminFaq := frt.router.Group("/admin/faqs", middleware.ValidateAdminJWT)
	adminFaq.POST("", frt.fc.HandleCreateFAQ)
	adminFaq.PUT("/:id", frt.fc.HandleUpdateFAQ)
	adminFaq.DELETE("/:id", frt.fc.HandleDeleteFAQ)
}

func (crt categoryRoutes) InitEndpoints() {
	category := crt.router.Group("/categories")
	category.GET("", crt.cc.HandleGetAllCategories)
	category.GET("/:id", crt.cc.HandleGetCategoryByID)

	adminCategory := crt.router.Group("/admin/categories", middleware.ValidateAdminJWT)
	adminCategory.POST("", crt.cc.HandleCreateCategory)
	adminCategory.PUT("/:id", crt.cc.HandleUpdateCategory)
	adminCategory.DELETE("/:is", crt.cc.HandleDeleteCategory)
}

func (prt pointRoutes) InitEndpoints() {
	point := prt.router.Group("/coins", middleware.ValidateJWT)
	point.GET("", prt.pc.HandleGetUserPoint)

	adminPoints := prt.router.Group("/admin/coins", middleware.ValidateAdminJWT)
	adminPoints.GET("", prt.pc.HandleGetAllPoint)
	adminPoints.GET("/:id", prt.pc.HandleGetPointByID)
}

func (rrt rewardRoutes) InitEndpoints() {
	reward := rrt.router.Group("/rewards")
	reward.GET("", rrt.rc.FindAllReward)
	reward.GET("/:id", rrt.rc.FindRewardById)

	adminReward := rrt.router.Group("/admin/rewards", middleware.ValidateAdminJWT)
	adminReward.POST("", rrt.rc.CreateReward)
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
	auth.POST("/forgot-password", art.ac.HandleForgotPassword)
	auth.POST("/new-password", art.ac.HandleNewPassword)

	token := auth.Group("/token")
	token.POST("/refresh", art.ac.HandleRefreshToken)

	otp := auth.Group("/otp")
	otp.POST("/validate", art.ac.HandleValidateOTP)
	otp.POST("/resend", art.ac.HandleRequestNewOTP)
}

func (urt userRoutes) InitEndpoints() {
	user := urt.router.Group("/users", middleware.ValidateJWT)
	user.GET("", urt.uc.HandleGetSelfUserData)
	user.PUT("/change-password", urt.uc.HandleChangePassword)
	user.PUT("", urt.uc.HandleUpdateData)

	admin := urt.router.Group("/admin/users", middleware.ValidateAdminJWT)
	admin.GET("", urt.uc.HandleGetAllUser)
	admin.GET("/:id", urt.uc.HandleGetUserByID)
	admin.PUT("/:id", urt.uc.HandleUpdateCustomerData)
	admin.DELETE("/:id", urt.uc.HandleDeleteCustomerData)
}

func (prt productRoutes) InitEndpoints() {
	product := prt.router.Group("/products")
	product.GET("", prt.pc.GetAll)
	product.GET("/:id", prt.pc.GetProductById)

	adminProduct := prt.router.Group("/admin/products", middleware.ValidateAdminJWT)
	adminProduct.POST("", prt.pc.InsertProduct)
	adminProduct.PUT("/:id", prt.pc.UpdateProduct)
	adminProduct.DELETE("/:id", prt.pc.DeleteProduct)
}

func (drt redeemRoutes) InitEndpoints() {
	adminRedeem := drt.router.Group("/admin/redeems", middleware.ValidateAdminJWT)
	adminRedeem.GET("", drt.dc.GetAllRedeem)
	adminRedeem.GET("/:id", drt.dc.AdminGetRedeemByID)
	adminRedeem.GET("/all", drt.dc.GetAllRedeemIncludeSoftDeleted)
	adminRedeem.PUT("/:id", drt.dc.UpdateRedeem)
	adminRedeem.DELETE("/:id", drt.dc.DeleteRedeem)

	redeem := drt.router.Group("/redeems", middleware.ValidateJWT)
	redeem.GET("", drt.dc.GetAllRedeemByUserID)
	redeem.GET("/:id", drt.dc.GetRedeemByID)
	redeem.POST("", drt.dc.CreateRedeem)
}
