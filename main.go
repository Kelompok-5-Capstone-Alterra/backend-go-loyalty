package main

import (
	internalMiddleware "backend-go-loyalty/internal/middleware"
	"backend-go-loyalty/pkg/bootstrapper"
	"backend-go-loyalty/pkg/config"
	"backend-go-loyalty/pkg/server"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UseCommonMiddlewares(router *echo.Echo) *echo.Echo {
	whitelist := config.GetWhitelistedURLS()
	router.Use(internalMiddleware.CorsMiddleware(whitelist))
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.Logger())
	return router
}

func main() {
	godotenv.Load(".env")
	router := config.InitRouter()
	router = UseCommonMiddlewares(router)
	env := config.GetEnvVariables()
	db := config.GetDatabase(env.DBAddress, env.DBUsername, env.DBPassword, env.DBName)
	// config.InitialMigration(db, &model.Role{}, &model.User{}, &model.OTP{}, &model.Product{}, &model.Reward{}) // Disabled because the app will use go migrate to declare DDL and other initial migration stuffs
	bootstrapper.InitEndpoints(router, db)
	server := server.NewServer(env.ServerAddress, router)
	server.ListenAndServe()
}
