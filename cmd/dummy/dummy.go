package main

import (
	"backend-go-loyalty/internal/model"
	"backend-go-loyalty/pkg/config"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".migration.env")
	env := config.GetEnvVariables()
	db := config.GetDatabase(env.DBAddress, env.DBUsername, env.DBPassword, env.DBName)
	switch os.Args[1] {
	case "coins":
		{
			err := db.Model(&model.UserCoin{}).Where("amount < ?", 10000).Update("amount", gorm.Expr("amount + ?", 10000)).Error
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
