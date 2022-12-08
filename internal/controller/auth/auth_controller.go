package authController

import (
	"backend-go-loyalty/internal/dto"
	authService "backend-go-loyalty/internal/service/auth"
	"backend-go-loyalty/pkg/response"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AuthController interface {
	HandleLogin(c echo.Context) error
	HandleSignUp(c echo.Context) error
	HandleValidateOTP(c echo.Context) error
	HandleRefreshToken(c echo.Context) error
	HandleRequestNewOTP(c echo.Context) error
}

type authController struct {
	as authService.AuthService
}

func NewAuthController(as authService.AuthService) authController {
	return authController{
		as: as,
	}
}

func (ac authController) HandleLogin(c echo.Context) error {
	req := dto.SignInRequest{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	data, err := ac.as.Login(c.Request().Context(), req)
	if err != nil {
		var code int
		if err.Error() == "record not found" || err.Error() == "wrong password" || err.Error() == "user not activated" {
			code = http.StatusUnauthorized
		} else {
			code = http.StatusInternalServerError
		}
		return response.ResponseError(code, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (ac authController) HandleSignUp(c echo.Context) error {
	req := dto.SignUpRequest{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	err = ac.as.SignUp(c.Request().Context(), req)
	if err != nil {
		var code int
		if err.Error() == "email already used" {
			code = http.StatusConflict
		} else {
			code = http.StatusInternalServerError
		}
		return response.ResponseError(code, err)
	}
	return response.ResponseSuccess(http.StatusCreated, echo.Map{"status": "PENDING_VERIFICATION"}, c)
}

func (ac authController) HandleValidateOTP(c echo.Context) error {
	req := dto.ValidateOTP{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	err = ac.as.ValidateOTP(c.Request().Context(), req)
	if err != nil {
		var code int
		if err.Error() == "invalid otp code" {
			code = http.StatusUnauthorized
		} else {
			code = http.StatusInternalServerError
		}
		return response.ResponseError(code, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{"status": "SUCCESS_ACTIVATED_USER"}, c)
}

func (ac authController) HandleRefreshToken(c echo.Context) error {
	rt := dto.RefreshRequest{}
	c.Bind(&rt)
	validate := validator.New()
	err := validate.Struct(&rt)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	data, err := ac.as.RegenerateToken(c.Request().Context(), rt.RefreshToken)
	if err != nil {
		var code int
		if err.Error() == "refresh token invalid" || err.Error() == "record not found" {
			code = http.StatusUnauthorized
		} else {
			code = http.StatusInternalServerError
		}
		return response.ResponseError(code, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}

func (ac authController) HandleRequestNewOTP(c echo.Context) error {
	req := dto.RequestNewOTP{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	err = ac.as.RequestNewOTP(c.Request().Context(), req.Email)
	if err != nil {
		var code int
		if err.Error() == "cannot send new otp to activated user" {
			code = http.StatusForbidden
		} else if err.Error() == "record not found" {
			code = http.StatusUnauthorized
		} else {
			code = http.StatusInternalServerError
		}
		return response.ResponseError(code, err)
	}
	return response.ResponseSuccess(http.StatusCreated, echo.Map{"status": "SUCCESS_SENT_OTP"}, c)
}
