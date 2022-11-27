package userController

import (
	"backend-go-loyalty/internal/dto"
	userService "backend-go-loyalty/internal/service/user"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	HandleChangePassword(c echo.Context) error
	HandleUpdateData(c echo.Context) error
}

type userController struct {
	us userService.UserServiceInterface
}

func NewUserController(us userService.UserServiceInterface) userController {
	return userController{
		us: us,
	}
}

func (uc userController) HandleChangePassword(c echo.Context) error {
	req := dto.UpdatePasswordRequest{}
	c.Bind(&req)
	id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			response.NewBaseResponse(http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				response.NewErrorResponseData(
					response.NewErrorResponseValue(
						"error",
						err.Error(),
					),
				),
				nil))
	}
	validate := validator.New()
	err = validate.Struct(&req)
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
	errs := uc.us.UpdatePassword(c.Request().Context(), req, id)
	if errs != nil {
		errVal := response.ErrorResponseValue{
			Key:   "error",
			Value: errs.Error(),
		}
		errRes := response.ErrorResponseData{errVal}
		return echo.NewHTTPError(http.StatusInternalServerError,
			response.NewBaseResponse(http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				errRes,
				nil))
	}
	return c.JSON(http.StatusOK, response.NewBaseResponse(http.StatusOK, http.StatusText(http.StatusOK), nil, echo.Map{
		"status": "SUCCESS_UPDATE_PASSWORD",
	}))
}

func (uc userController) HandleUpdateData(c echo.Context) error {
	req := dto.UserUpdate{}
	c.Bind(&req)
	id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			response.NewBaseResponse(http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
				response.NewErrorResponseData(
					response.NewErrorResponseValue(
						"error",
						err.Error(),
					),
				),
				nil))
	}
	validate := validator.New()
	err = validate.Struct(&req)
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
	data, err := uc.us.UpdateUserData(c.Request().Context(), req, id)
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
