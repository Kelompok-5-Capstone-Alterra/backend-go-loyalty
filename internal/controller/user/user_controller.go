package userController

import (
	"backend-go-loyalty/internal/dto"
	userService "backend-go-loyalty/internal/service/user"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	HandleChangePassword(c echo.Context) error
	HandleUpdateData(c echo.Context) error
	HandleUpdateCustomerData(c echo.Context) error
	HandleDeleteCustomerData(c echo.Context) error
}

type userController struct {
	us userService.UserServiceInterface
}

func NewUserController(us userService.UserServiceInterface) userController {
	return userController{
		us: us,
	}
}

func (uc userController) HandleUpdateCustomerData(c echo.Context) error {
	req := dto.UserUpdate{}
	c.Bind(&req)
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorSingle(err, http.StatusBadRequest)
	}
	_, err = uc.us.UpdateUserData(c.Request().Context(), req, id)
	if err != nil {
		if err.Error() == "record not found" {
			return responseErrorSingle(err, http.StatusNoContent)
		}
		return responseErrorSingle(err, http.StatusInternalServerError)
	}
	return responseSuccess(c, http.StatusOK, echo.Map{
		"status": "SUCCESS_UPDATED_USER",
	})
}

func (uc userController) HandleDeleteCustomerData(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorSingle(err, http.StatusBadRequest)
	}
	err = uc.us.DeleteUserData(c.Request().Context(), id)
	if err != nil {
		return responseErrorSingle(err, http.StatusInternalServerError)
	}
	return responseSuccess(c, http.StatusOK, echo.Map{
		"status": "SUCCESS_DELETED_USER",
	})
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

func responseErrorSingle(err error, code int) error {
	return echo.NewHTTPError(code,
		response.NewBaseResponse(code,
			http.StatusText(code),
			response.NewErrorResponseData(
				response.NewErrorResponseValue(
					"error",
					err.Error(),
				),
			),
			nil))
}

func responseSuccess(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, response.NewBaseResponse(code, http.StatusText(code), nil, data))
}
