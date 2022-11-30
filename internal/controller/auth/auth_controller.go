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
		errRes := response.ErrorResponseData{}
		for _, val := range err.(validator.ValidationErrors) {
			var errVal response.ErrorResponseValue
			errVal.Key = val.StructField()
			errVal.Value = val.Tag()
			errRes = append(errRes, errVal)
		}
		return echo.NewHTTPError(http.StatusBadRequest,
			response.NewBaseResponse(http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				errRes,
				nil))
	}
	data, err := ac.as.Login(c.Request().Context(), req)
	if err != nil {
		errVal := response.ErrorResponseValue{
			Key:   "error",
			Value: err.Error(),
		}
		errRes := response.ErrorResponseData{errVal}
		return echo.NewHTTPError(http.StatusInternalServerError,
			response.NewBaseResponse(http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				errRes,
				nil))
	}
	return c.JSON(http.StatusOK, response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, data))
}
func (ac authController) HandleSignUp(c echo.Context) error {
	req := dto.SignUpRequest{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		errRes := response.ErrorResponseData{}
		for _, val := range err.(validator.ValidationErrors) {
			var errVal response.ErrorResponseValue
			errVal.Key = val.StructField()
			errVal.Value = val.Tag()
			errRes = append(errRes, errVal)
		}
		return echo.NewHTTPError(http.StatusBadRequest,
			response.NewBaseResponse(http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				errRes,
				nil))
	}
	err = ac.as.SignUp(c.Request().Context(), req)
	if err != nil {
		errVal := response.ErrorResponseValue{
			Key:   "error",
			Value: err.Error(),
		}
		errRes := response.ErrorResponseData{errVal}
		return echo.NewHTTPError(http.StatusInternalServerError,
			response.NewBaseResponse(http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				errRes,
				nil))
	}
	return c.JSON(http.StatusOK, response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, echo.Map{
		"status": "PENDING_VERIFICATION",
	}))
}

func (ac authController) HandleValidateOTP(c echo.Context) error {
	req := dto.ValidateOTP{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		errRes := response.ErrorResponseData{}
		for _, val := range err.(validator.ValidationErrors) {
			var errVal response.ErrorResponseValue
			errVal.Key = val.StructField()
			errVal.Value = val.Tag()
			errRes = append(errRes, errVal)
		}
		return echo.NewHTTPError(http.StatusBadRequest,
			response.NewBaseResponse(http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				errRes,
				nil))
	}
	err = ac.as.ValidateOTP(c.Request().Context(), req)
	if err != nil {
		errVal := response.ErrorResponseValue{
			Key:   "error",
			Value: err.Error(),
		}
		errRes := response.ErrorResponseData{errVal}
		return echo.NewHTTPError(http.StatusInternalServerError,
			response.NewBaseResponse(http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				errRes,
				nil))
	}
	return c.JSON(http.StatusOK, response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, echo.Map{
		"status": "SUCCESS_ACTIVATED_USER",
	}))
}

func (ac authController) HandleRefreshToken(c echo.Context) error {
	rt := dto.RefreshRequest{}
	c.Bind(&rt)
	validate := validator.New()
	err := validate.Struct(&rt)
	if err != nil {
		errRes := response.ErrorResponseData{}
		for _, val := range err.(validator.ValidationErrors) {
			var errVal response.ErrorResponseValue
			errVal.Key = val.StructField()
			errVal.Value = val.Tag()
			errRes = append(errRes, errVal)
		}
		return echo.NewHTTPError(http.StatusBadRequest,
			response.NewBaseResponse(http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				errRes,
				nil))
	}
	data, err := ac.as.RegenerateToken(c.Request().Context(), rt.RefreshToken)
	if err != nil {
		errVal := response.ErrorResponseValue{
			Key:   "error",
			Value: err.Error(),
		}
		errRes := response.ErrorResponseData{errVal}
		return echo.NewHTTPError(http.StatusInternalServerError,
			response.NewBaseResponse(http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				errRes,
				nil))
	}
	return c.JSON(http.StatusOK, response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, data))
}

func (ac authController) HandleRequestNewOTP(c echo.Context) error {
	req := dto.RequestNewOTP{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(&req)
	if err != nil {
		errRes := response.ErrorResponseData{}
		for _, val := range err.(validator.ValidationErrors) {
			var errVal response.ErrorResponseValue
			errVal.Key = val.StructField()
			errVal.Value = val.Tag()
			errRes = append(errRes, errVal)
		}
		return echo.NewHTTPError(http.StatusBadRequest,
			response.NewBaseResponse(http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				errRes,
				nil))
	}
	err = ac.as.RequestNewOTP(c.Request().Context(), req.Email)
	if err != nil {
		errVal := response.ErrorResponseValue{
			Key:   "error",
			Value: err.Error(),
		}
		errRes := response.ErrorResponseData{errVal}
		return echo.NewHTTPError(http.StatusInternalServerError,
			response.NewBaseResponse(http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				errRes,
				nil))
	}
	return c.JSON(http.StatusOK, response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, echo.Map{
		"status": "SUCCESS_SENT_OTP",
	}))
}
