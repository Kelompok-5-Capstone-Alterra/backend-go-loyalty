package routes

import (
	authController "backend-go-loyalty/internal/controller/auth"
	categoryController "backend-go-loyalty/internal/controller/category"
	faqController "backend-go-loyalty/internal/controller/faq"
	paymentController "backend-go-loyalty/internal/controller/payment"
	pingController "backend-go-loyalty/internal/controller/ping"
	pointController "backend-go-loyalty/internal/controller/point"
	productController "backend-go-loyalty/internal/controller/product"
	redeemController "backend-go-loyalty/internal/controller/redeem"
	rewardController "backend-go-loyalty/internal/controller/reward"
	transactionController "backend-go-loyalty/internal/controller/transaction"
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

type transactionRoutes struct {
	tc     transactionController.ITransactionController
	router *echo.Echo
}

type paymentRoutes struct {
	pc     paymentController.IPaymentController
	router *echo.Echo
}

func NewPaymentRoutes(pc paymentController.IPaymentController, router *echo.Echo) paymentRoutes {
	return paymentRoutes{
		pc:     pc,
		router: router,
	}
}

func NewTransactionRoutes(tc transactionController.ITransactionController, router *echo.Echo) transactionRoutes {
	return transactionRoutes{
		tc:     tc,
		router: router,
	}
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

func (prt paymentRoutes) InitEndpoints() {
	prt.router.POST("/payment/webhook", prt.pc.HandleNotification, middleware.ValidateXenditCallback)
}

func (trt transactionRoutes) InitEndpoints() {
	trt.router.GET("/admin/transactions", trt.tc.HandleGetAllTransaction, middleware.ValidateAdminJWT)
	trt.router.GET("/admin/transactions/:id", trt.tc.HandleGetTransactionByID, middleware.ValidateAdminJWT)

	trt.router.GET("/transactions", trt.tc.HandleGetTransactionsByUserID, middleware.ValidateJWT)
	trt.router.GET("/transactions/:id", trt.tc.HandleGetTransactionByIDByUser, middleware.ValidateJWT)
	trt.router.GET("/transactions", trt.tc.HandleCreateTransaction, middleware.ValidateJWT)
}

func (frt faqRoutes) InitEndpoints() {
	frt.router.GET("/faqs", frt.fc.HandleGetAllFAQByKeyword)
	frt.router.GET("/faqs/:id", frt.fc.HandleGetFAQByID)
	frt.router.POST("/admin/faqs", frt.fc.HandleCreateFAQ, middleware.ValidateAdminJWT)
	frt.router.PUT("/admin/faqs/:id", frt.fc.HandleUpdateFAQ, middleware.ValidateAdminJWT)
	frt.router.DELETE("/admin/faqs/:id", frt.fc.HandleDeleteFAQ, middleware.ValidateAdminJWT)
}

func (crt categoryRoutes) InitEndpoints() {
	crt.router.GET("/categories", crt.cc.HandleGetAllCategories)
	crt.router.GET("/categories/:id", crt.cc.HandleGetCategoryByID)
	crt.router.POST("/categories", crt.cc.HandleCreateCategory, middleware.ValidateAdminJWT)
	crt.router.PUT("/categories/:id", crt.cc.HandleUpdateCategory, middleware.ValidateAdminJWT)
	crt.router.DELETE("/categories/:id", crt.cc.HandleDeleteCategory, middleware.ValidateAdminJWT)
}

func (prt pointRoutes) InitEndpoints() {
	prt.router.GET("/coins", prt.pc.HandleGetUserPoint, middleware.ValidateJWT)
	prt.router.GET("/admin/coins", prt.pc.HandleGetAllPoint, middleware.ValidateAdminJWT)
	prt.router.GET("/admin/coins/:id", prt.pc.HandleGetPointByID, middleware.ValidateAdminJWT)
}

func (rrt rewardRoutes) InitEndpoints() {
	rrt.router.GET("/rewards", rrt.rc.FindAllReward)
	rrt.router.GET("/rewards/:id", rrt.rc.FindRewardById)
	rrt.router.POST("/admin/rewards", rrt.rc.CreateReward, middleware.ValidateAdminJWT)
	rrt.router.PUT("/admin/rewards/:id", rrt.rc.UpdateReward, middleware.ValidateAdminJWT)
	rrt.router.DELETE("/admin/rewards/:id", rrt.rc.DeleteReward, middleware.ValidateAdminJWT)
}

func (prt pingRoutes) InitEndpoints() {
	ping := prt.router.Group("/ping")
	ping.GET("", prt.pc.HandlePing)
}

func (art authRoutes) InitEndpoints() {
	art.router.POST("/auth/signin", art.ac.HandleLogin)
	art.router.POST("/auth/signup", art.ac.HandleSignUp)
	art.router.POST("/auth/forgot-password", art.ac.HandleForgotPassword)
	art.router.POST("/auth/new-password", art.ac.HandleNewPassword)
	art.router.POST("/auth/token/refresh", art.ac.HandleRefreshToken)
	art.router.POST("/auth/otp/validate", art.ac.HandleValidateOTP)
	art.router.POST("/auth/otp/resend", art.ac.HandleRequestNewOTP)
}

func (urt userRoutes) InitEndpoints() {
	urt.router.GET("/users", urt.uc.HandleGetSelfUserData, middleware.ValidateJWT)
	urt.router.PUT("/users/change-password", urt.uc.HandleChangePassword, middleware.ValidateJWT)
	urt.router.PUT("/users", urt.uc.HandleUpdateData, middleware.ValidateJWT)
	urt.router.GET("/admin/users", urt.uc.HandleGetAllUser, middleware.ValidateAdminJWT)
	urt.router.GET("/admin/users/:id", urt.uc.HandleGetUserByID, middleware.ValidateAdminJWT)
	urt.router.PUT("/admin/users/:id", urt.uc.HandleUpdateCustomerData, middleware.ValidateAdminJWT)
	urt.router.DELETE("/admin/users/:id", urt.uc.HandleDeleteCustomerData, middleware.ValidateAdminJWT)
}

func (prt productRoutes) InitEndpoints() {
	prt.router.GET("/products", prt.pc.GetAll)
	prt.router.GET("/products/:id", prt.pc.GetProductById)
	prt.router.POST("/admin/products", prt.pc.InsertProduct, middleware.ValidateAdminJWT)
	prt.router.PUT("/admin/products/:id", prt.pc.UpdateProduct, middleware.ValidateAdminJWT)
	prt.router.DELETE("/admin/products/:id", prt.pc.DeleteProduct, middleware.ValidateAdminJWT)
}

func (drt redeemRoutes) InitEndpoints() {
	drt.router.GET("/admin/redeems", drt.dc.GetAllRedeem, middleware.ValidateAdminJWT)
	drt.router.GET("/admin/redeems/:id", drt.dc.AdminGetRedeemByID, middleware.ValidateAdminJWT)
	drt.router.GET("/admin/redeems/all", drt.dc.GetAllRedeemIncludeSoftDeleted, middleware.ValidateAdminJWT)
	drt.router.PUT("/admin/redeems/:id", drt.dc.UpdateRedeem, middleware.ValidateAdminJWT)
	drt.router.DELETE("/admin/redeems/:id", drt.dc.DeleteRedeem, middleware.ValidateAdminJWT)
	drt.router.GET("/redeems", drt.dc.GetAllRedeemByUserID, middleware.ValidateJWT)
	drt.router.GET("/redeems/:id", drt.dc.GetRedeemByID, middleware.ValidateJWT)
	drt.router.POST("/redeems", drt.dc.CreateRedeem, middleware.ValidateJWT)
}
