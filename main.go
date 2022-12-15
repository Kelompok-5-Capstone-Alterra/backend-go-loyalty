package main

import (
	internalMiddleware "backend-go-loyalty/internal/middleware"
	"backend-go-loyalty/pkg/bootstrapper"
	"backend-go-loyalty/pkg/config"
	"backend-go-loyalty/pkg/server"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xendit/xendit-go/client"
)

func UseCommonMiddlewares(router *echo.Echo) *echo.Echo {
	whitelist := config.GetWhitelistedURLS()
	router.Use(internalMiddleware.CorsMiddleware(whitelist))
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(internalMiddleware.MiddlewareLogging)
	// router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
	// 	Output: router.Logger.Output(),
	// }))
	return router
}

func main() {
	godotenv.Load(".env")
	router := config.InitRouter()
	router = UseCommonMiddlewares(router)
	env := config.GetEnvVariables()
	db := config.GetDatabase(env.DBAddress, env.DBUsername, env.DBPassword, env.DBName)
	xenditClient := client.New(env.XenditServerKey)
	// config.InitialMigration(db, &model.Role{}, &model.User{}, &model.OTP{}, &model.Product{}, &model.Reward{}) // Disabled because the app will use go migrate to declare DDL and other initial migration stuffs
	bootstrapper.InitEndpoints(router, db, xenditClient.Invoice)
	server := server.NewServer(env.ServerAddress, router)
	server.ListenAndServe()
}
