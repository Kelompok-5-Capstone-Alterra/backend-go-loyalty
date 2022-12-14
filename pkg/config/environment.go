package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Env struct {
	ServerAddress   string
	DBAddress       string
	DBUsername      string
	DBPassword      string
	DBName          string
	XenditServerKey string
}

func GetEnvVariables() Env {
	return Env{
		ServerAddress:   os.Getenv("SERVER_ADDRESS"),
		DBAddress:       os.Getenv("DB_ADDRESS"),
		DBUsername:      os.Getenv("DB_USERNAME"),
		DBPassword:      os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		XenditServerKey: os.Getenv("XENDIT_SERVER_KEY"),
	}
}

func GetJWTKey() string {
	return os.Getenv("JWT_KEY")
}

type TokenEnv struct {
	AccessTokenTTLHour  int64
	RefreshTokenTTLHour int64
}

func GetTokenEnv() TokenEnv {
	tokenTTL, err := strconv.ParseInt(os.Getenv("ACCESS_TOKEN_TTL_HOUR"), 10, 64)
	if err != nil {
		log.Println(err.Error())
	}
	refreshTTL, err := strconv.ParseInt(os.Getenv("REFRESH_TOKEN_TTL_HOUR"), 10, 64)
	if err != nil {
		log.Println(err.Error())
	}
	return TokenEnv{
		AccessTokenTTLHour:  tokenTTL,
		RefreshTokenTTLHour: refreshTTL,
	}
}

func GetWhitelistedURLS() map[string]bool {
	whitelist := os.Getenv("WHITELISTED_URLS")
	whitelistArray := strings.Split(whitelist, ",")
	urls := make(map[string]bool)
	for _, value := range whitelistArray {
		urls[value] = true
	}
	return urls
}
