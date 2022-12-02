package main

import (
	"backend-go-loyalty/internal/middleware"
	"backend-go-loyalty/internal/model"
	"backend-go-loyalty/pkg/bootstrapper"
	"backend-go-loyalty/pkg/config"
	"backend-go-loyalty/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	router := config.InitRouter()
	whitelist := config.GetWhitelistedURLS()
	router.Use(middleware.CorsMiddleware(whitelist))
	env := config.GetEnvVariables()
	db := config.GetDatabase(env.DBAddress, env.DBUsername, env.DBPassword, env.DBName)
	config.InitialMigration(db, &model.Role{}, &model.User{}, &model.OTP{}, &model.Product{}, &model.Reward{})
	bootstrapper.InitEndpoints(router, db)
	server := server.NewServer(env.ServerAddress, router)
	server.ListenAndServe()
}
