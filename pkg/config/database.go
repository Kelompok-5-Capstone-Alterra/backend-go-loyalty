package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase(address, username, password, name string) *gorm.DB {
	log.Println("INFO GetDatabase: starting database connection processes")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, name)

	var db *gorm.DB
	var err error
	for {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			// Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Println("INFO GetDatabase: re-attempting to connect to database...")
			time.Sleep(1 * time.Second)
			continue
		} else {
			break
		}

	}
	log.Println("INFO GetDatabase: Successfully established connection with ", dsn)
	return db
}

func InitialMigration(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}

func InitDatabase(envVariables Env) *gorm.DB {
	return GetDatabase(
		envVariables.DBAddress,
		envVariables.DBUsername,
		envVariables.DBPassword,
		envVariables.DBName,
	)
}
