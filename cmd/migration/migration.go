package main

import (
	"backend-go-loyalty/pkg/config"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	godotenv.Load(".migration.env")
	env := config.GetEnvVariables()
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s", env.DBUsername, env.DBPassword, env.DBAddress, env.DBName)
	m, err := migrate.New("file://./database/migration", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		recover := recover()
		if recover != nil {
			fmt.Println(recover.(error).Error())
		}
	}()
	switch os.Args[1] {
	case "up":
		{
			log.Println("Migrating Database Up...")
			err := m.Up()
			fmt.Printf("migration(up): Error eccoured -> %v\n",err.Error())
		}
	case "down":
		{
			log.Println("Migrating Database Down...")
			err := m.Down()
			if err != nil{
				fmt.Printf("migration(down): Error eccoured -> %v\n",err.Error())
				return
			}
		}
	default:
		{
			fmt.Println("You've entered wrong command")
		}
	}

}
