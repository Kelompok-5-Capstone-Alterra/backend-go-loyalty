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
		return responseErrorRequestBody(http.StatusBadRequest,err)
	}
	data, err := ac.as.Login(c.Request().Context(), req)
	if err != nil {
		var code int
		if err.Error() == "record not found" || err.Error() == "wrong password" || err.Error() == "user not activated"{
			code = http.StatusUnauthorized
		} else{
			code = http.StatusInternalServerError
		}
		return responseError(code,err)
	}
	return responseSuccess(http.StatusOK,data,c)
}
func (ac authController) HandleSignUp(c echo.Context) error {
	req := dto.SignUpRequest{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		return responseErrorRequestBody(http.StatusBadRequest,err)
	}
	err = ac.as.SignUp(c.Request().Context(), req)
	if err != nil {
		var code int
		if err.Error() == "email alreasy used"{
			code = http.StatusConflict
		} else{
			code = http.StatusInternalServerError
		}
		return responseError(code, err)
	}
	return responseSuccess(http.StatusOK,echo.Map{"status":"PENDING_VERIFICATION"},c)
}

func (ac authController) HandleValidateOTP(c echo.Context) error {
	req := dto.ValidateOTP{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		return responseErrorRequestBody(http.StatusBadRequest,err)
	}
	err = ac.as.ValidateOTP(c.Request().Context(), req)
	if err != nil {
		var code int
		if err.Error() == "invalid otp code"{
			code = http.StatusUnauthorized
		} else{
			code = http.StatusInternalServerError
		}
		return responseError(code, err)
	}
	return responseSuccess(http.StatusOK,echo.Map{"status": "SUCCESS_ACTIVATED_USER"},c)
}

func (ac authController) HandleRefreshToken(c echo.Context) error {
	rt := dto.RefreshRequest{}
	c.Bind(&rt)
	validate := validator.New()
	err := validate.Struct(&rt)
	if err != nil {
		return responseErrorRequestBody(http.StatusBadRequest,err)
	}
	data, err := ac.as.RegenerateToken(c.Request().Context(), rt.RefreshToken)
	if err != nil {
		var code int
		if err.Error() == "refresh token invalid"{
			code = http.StatusUnauthorized
		} else{
			code = http.StatusInternalServerError
		}
		return responseError(code,err)
	}
	return responseSuccess(http.StatusOK,data,c)
}

func (ac authController) HandleRequestNewOTP(c echo.Context) error {
	req := dto.RequestNewOTP{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		return responseErrorRequestBody(http.StatusBadRequest,err)
	}
	err = ac.as.RequestNewOTP(c.Request().Context(), req.Email)
	if err != nil {
		var code int
		if err.Error() == "cannot send new otp to activated user"{
			code = http.StatusForbidden
		} else{
			code = http.StatusInternalServerError
		}
		return responseError(code,err)
	}
	return responseSuccess(http.StatusOK,echo.Map{"status": "SUCCESS_SENT_OTP"},c)
}


func responseErrorRequestBody(code int, err error)error{
	errRes := response.ErrorResponseData{}
		for _, val := range err.(validator.ValidationErrors) {
			var errVal response.ErrorResponseValue
			errVal.Key = val.StructField()
			errVal.Value = val.Tag()
			errRes = append(errRes, errVal)
		}
		return echo.NewHTTPError(code,
			response.NewBaseResponse(code,
				http.StatusText(code),
				errRes,
				nil))
}


func responseError(code int, err error) error{
	errVal := response.ErrorResponseValue{
		Key:   "error",
		Value: err.Error(),
	}
	errRes := response.ErrorResponseData{errVal}
	return echo.NewHTTPError(code,
		response.NewBaseResponse(code,
			http.StatusText(code),
			errRes,
			nil))
}

func responseSuccess(code int,data interface{}, c echo.Context) error{
	return c.JSON(code, response.NewBaseResponse(code, http.StatusText(code), nil, data))
}