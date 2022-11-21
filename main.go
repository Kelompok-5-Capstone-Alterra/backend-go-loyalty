package main

import (
	"backend-go-loyalty/pkg/bootstrapper"
	"backend-go-loyalty/pkg/config"
	"backend-go-loyalty/pkg/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	router := config.InitRouter()
	env := config.GetEnvVariables()
	bootstrapper.InitEndpoints(router)
	server := server.NewServer(env.ServerAddress, router)
	server.ListenAndServe()
}
