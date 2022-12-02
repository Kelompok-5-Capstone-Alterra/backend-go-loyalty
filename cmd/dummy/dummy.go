package main

import (
	"backend-go-loyalty/pkg/config"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	env := config.GetEnvVariables()
	db := config.GetDatabase(env.DBAddress, env.DBUsername, env.DBPassword, env.DBName)
	switch os.Args[1] {
	case "coins":
		{
			// db.Model(&model.)
			err := db.Raw(`
			UPDATE user_coins 
			SET amount = amount + 10000
			`).Error
			if err != nil {
				return
			}
		}
	default:
		{
			fmt.Println("invalid command")
		}
	}
}
