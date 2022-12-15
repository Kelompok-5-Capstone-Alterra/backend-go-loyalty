package middleware

import (
	"backend-go-loyalty/pkg/config"
	"backend-go-loyalty/pkg/response"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func ValidateAdminJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
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
				return echo.NewHTTPError(http.StatusUnauthorized,
					response.NewBaseResponse(http.StatusUnauthorized,
						http.StatusText(http.StatusUnauthorized),
						response.NewErrorResponseData(
							response.NewErrorResponseValue("error", err.Error()),
						), nil))
			}

			if token.Valid {
				data := claims["data"].(map[string]interface{})
				if data["role"].(map[string]interface{})["id"].(float64) == float64(config.ADMIN_ROLE_ID) {
					return next(c)
				}
				return echo.NewHTTPError(http.StatusUnauthorized,
					response.NewBaseResponse(http.StatusUnauthorized,
						http.StatusText(http.StatusUnauthorized),
						response.NewErrorResponseData(
							response.NewErrorResponseValue("error", "unauthorized"),
						), nil))
			}
			return nil
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized,
				response.NewBaseResponse(http.StatusUnauthorized,
					http.StatusText(http.StatusUnauthorized),
					response.NewErrorResponseData(
						response.NewErrorResponseValue("error", "unauthorized"),
					), nil))
		}
	}
}

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["Authorization"] != nil {
			auth := strings.Split(c.Request().Header["Authorization"][0], " ")
			token, err := jwt.Parse(auth[1], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, errors.New("unauthorized")
				}
				return []byte(config.GetJWTKey()), nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized,
					response.NewBaseResponse(http.StatusUnauthorized,
						http.StatusText(http.StatusUnauthorized),
						response.NewErrorResponseData(
							response.NewErrorResponseValue("error", err.Error()),
						), nil))
			}

			if token.Valid {
				return next(c)
			}
			return nil
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized,
				response.NewBaseResponse(http.StatusUnauthorized,
					http.StatusText(http.StatusUnauthorized),
					response.NewErrorResponseData(
						response.NewErrorResponseValue("error", "unauthorized"),
					), nil))
		}
	}
}

func CorsMiddleware(whitelistedUrls map[string]bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestOriginUrl := c.Request().Header.Get("Origin")
			if whitelistedUrls[requestOriginUrl] {
				c.Response().Header().Set("Access-Control-Allow-Origin", requestOriginUrl)
			}
			c.Response().Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			log.Printf("INFO CorsMiddleware: received request from %s %v", requestOriginUrl, whitelistedUrls[requestOriginUrl])
			if c.Request().Method != http.MethodOptions {
				return next(c)
			}
			return nil
		}
	}
}

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("Incoming Request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}

func makeLogEntry(c echo.Context) *logrus.Entry {
	if c == nil {
		return logrus.WithFields(logrus.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return logrus.WithFields(logrus.Fields{
		"At":     time.Now().Format("2006-01-02 15:04:05"),
		"Method": c.Request().Method,
		"URI":    c.Request().URL.String(),
		"IP":     c.Request().RemoteAddr,
	})
}

func ValidateXenditCallback(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := os.Getenv("X_CALLBACK_TOKEN")
		if c.Request().Header["X-Callback-Token"][0] == token {
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
				"code":  http.StatusUnauthorized,
				"error": "ACCESS_DENIED",
			})
		}
	}
}
