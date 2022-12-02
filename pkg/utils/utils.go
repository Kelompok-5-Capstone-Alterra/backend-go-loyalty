package utils

import (
	"backend-go-loyalty/internal/dto"
	"backend-go-loyalty/pkg/config"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/xlzd/gotp"
	"gopkg.in/gomail.v2"
)

func HashPassword(password string) string {
	hash := sha256.New()
	modifiedPass := fmt.Sprint(password, os.Getenv("PASSWORD_HASH_KEY"))
	hash.Write([]byte(modifiedPass))
	passHash := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return passHash
}

func CreateLoginToken(userID uuid.UUID, data dto.JWTData) (string, string) {
	claims := jwt.MapClaims{}
	claims["sub"] = userID
	claims["data"] = data
	claims["created_at"] = time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, _ := token.SignedString([]byte(config.GetJWTKey()))

	claims = jwt.MapClaims{}
	claims["sub"] = userID
	claims["created_at"] = time.Now()
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, _ := rtoken.SignedString([]byte(config.GetJWTKey()))

	return accessToken, refreshToken
}

func GenerateOTP() string {
	length := 4
	str := gotp.RandomSecret(length)
	str = str[:len(str)-1]
	return str
}

type SMTPConfig struct {
	Host        string
	Port        int
	Sender      string
	Email       string
	AppPassword string
}

func GetSMTPConfig() SMTPConfig {
	port, _ := strconv.ParseInt(os.Getenv("CONFIG_SMTP_PORT"), 10, 64)
	return SMTPConfig{
		Host:        os.Getenv("CONFIG_SMTP_HOST"),
		Port:        int(port),
		Sender:      os.Getenv("CONFIG_SENDER_NAME"),
		Email:       os.Getenv("CONFIG_AUTH_EMAIL"),
		AppPassword: os.Getenv("CONFIG_AUTH_PASSWORD"),
	}
}

func SendOTPToEmail(otp string, target string) error {
	env := GetSMTPConfig()
	mailBody := fmt.Sprint("OTP: ", otp)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", env.Sender)
	mailer.SetHeader("To", target)
	mailer.SetHeader("Subject", "Digital Outlet Account Verification")
	mailer.SetBody("text/html", mailBody)

	dialer := gomail.NewDialer(
		env.Host,
		env.Port,
		env.Email,
		env.AppPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	log.Println("Mail sent!")
	return nil
}

func GetDataFromRefreshToken(rt string) (uuid.UUID, time.Time, error) {
	jwtKey := config.GetJWTKey()
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(rt, claims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return uuid.UUID{}, time.Time{}, err
	}

	if token.Valid {
		res, err := time.Parse(time.RFC3339, claims["created_at"].(string))
		if err != nil {
			return uuid.UUID{}, time.Time{}, err
		}
		env := config.GetTokenEnv()
		if time.Now().Sub(res).Hours() > float64(env.RefreshTokenTTLHour) {
			return uuid.UUID{}, time.Time{}, errors.New("token expired")
		}
		idStr := fmt.Sprintf("%v", claims["sub"])
		idConv, err := uuid.Parse(idStr)
		if err != nil {
			return uuid.UUID{}, time.Time{}, err
		}
		return idConv, res, nil
	}
	return uuid.UUID{}, time.Time{}, errors.New("token invalid")
}

func GetUserIDFromJWT(c echo.Context) (uuid.UUID, error) {
	if c.Request().Header["Authorization"] != nil {
		claims := jwt.MapClaims{}
		auth := strings.Split(c.Request().Header["Authorization"][0], " ")
		token, err := jwt.ParseWithClaims(auth[1], claims, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unauthorized")
			}
			return []byte(config.GetJWTKey()), nil
		})
		if err != nil {
			return uuid.UUID{}, err
		}

		if token.Valid {
			idStr := fmt.Sprintf("%v", claims["sub"])
			idConv, err := uuid.Parse(idStr)
			return idConv, err
		}
		return uuid.UUID{}, errors.New("unauthorized")
	} else {
		return uuid.UUID{}, errors.New("unauthorized")
	}
}
