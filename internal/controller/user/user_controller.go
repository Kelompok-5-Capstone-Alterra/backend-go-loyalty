package userController

import (
	"backend-go-loyalty/internal/dto"
	userService "backend-go-loyalty/internal/service/user"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserControllerInterface interface {
	HandleChangePassword(c echo.Context) error
	HandleUpdateData(c echo.Context) error
	HandleUpdateCustomerData(c echo.Context) error
	HandleDeleteCustomerData(c echo.Context) error
	HandleGetAllUser(c echo.Context) error
	HandleGetUserByID(c echo.Context) error
	HandleGetSelfUserData(c echo.Context) error
}

type userController struct {
	us userService.UserServiceInterface
}

func NewUserController(us userService.UserServiceInterface) userController {
	return userController{
		us: us,
	}
}

func (uc userController) HandleGetSelfUserData(c echo.Context) error {
	id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := uc.us.GetUserByID(c.Request().Context(), id)
	if err != nil {
		var code int
		if err.Error() == "record not found" {
			code = http.StatusUnauthorized
		} else {
			code = http.StatusInternalServerError
		}
		return response.ResponseError(code, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}

func (uc userController) HandleGetAllUser(c echo.Context) error {
	query := c.QueryParam("name")
	data, err := uc.us.GetUsers(c.Request().Context(), query)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	if len(data) == 0 {
		return response.ResponseSuccess(http.StatusNoContent, nil, c)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (uc userController) HandleGetUserByID(c echo.Context) error {
	param := c.Param("id")
	id, err := uuid.Parse(param)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := uc.us.GetUserByID(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "record not found" {
			return response.ResponseSuccess(http.StatusNoContent, nil, c)
		}
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}

func (uc userController) HandleUpdateCustomerData(c echo.Context) error {
	req := dto.UserUpdate{}
	c.Bind(&req)
	param := c.Param("id")
	id, err := uuid.Parse(param)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	_, err = uc.us.UpdateUserData(c.Request().Context(), req, id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_UPDATED_USER",
	}, c)
}

func (uc userController) HandleDeleteCustomerData(c echo.Context) error {
	param := c.Param("id")
	id, err := uuid.Parse(param)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	err = uc.us.DeleteUserData(c.Request().Context(), id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_DELETED_USER",
	}, c)
}

func (uc userController) HandleChangePassword(c echo.Context) error {
	req := dto.UpdatePasswordRequest{}
	c.Bind(&req)
	id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	validate := validator.New()
	err = validate.Struct(&req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	errs := uc.us.UpdatePassword(c.Request().Context(), req, id)
	if errs != nil {
		var code int
		if err.Error() == "new password must be different from old password" {
			code = http.StatusBadRequest
		} else if err.Error() == "password not match" {
			code = http.StatusUnauthorized
		} else {
			code = http.StatusInternalServerError
		}
		return response.ResponseError(code, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_UPDATE_PASSWORD",
	}, c)
}

func (uc userController) HandleUpdateData(c echo.Context) error {
	req := dto.UserUpdate{}
	c.Bind(&req)
	id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	validate := validator.New()
	err = validate.Struct(&req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	data, err := uc.us.UpdateUserData(c.Request().Context(), req, id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
