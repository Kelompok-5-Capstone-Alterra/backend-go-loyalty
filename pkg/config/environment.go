package config

import "os"

type Env struct {
	ServerAddress string
	DBAddress     string
	DBUsername    string
	DBPassword    string
	DBName        string
}

func GetEnvVariables() Env {
	return Env{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
		DBAddress:     os.Getenv("DB_ADDRESS"),
		DBUsername:    os.Getenv("DB_USERNAME"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
	}
}
