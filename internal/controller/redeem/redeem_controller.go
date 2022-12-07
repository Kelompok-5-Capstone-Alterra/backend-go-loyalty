package redeemController

import (
	"backend-go-loyalty/internal/dto"
	redeemService "backend-go-loyalty/internal/service/redeem"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type IRedeemController interface {
	CreateRedeem(c echo.Context) error
	GetAllRedeemByUserID(c echo.Context) error
	GetAllRedeem(c echo.Context) error
	GetAllRedeemIncludeSoftDeleted(c echo.Context) error
	GetRedeemByID(c echo.Context) error
	UpdateRedeem(c echo.Context) error
	DeleteRedeem(c echo.Context) error
}

type redeemController struct {
	ds redeemService.IRedeemService
}

func NewRedeemController(ds redeemService.IRedeemService) redeemController {
	return redeemController{
		ds: ds,
	}
}

func (dc redeemController) GetAllRedeem(c echo.Context) error {
	data, err := dc.ds.GetAllRedeems(c.Request().Context())
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	if len(data) == 0 {
		return response.ResponseSuccess(http.StatusNoContent, data, c)
	} else {
		return response.ResponseSuccess(http.StatusOK, data, c)
	}
}

func (dc redeemController) GetAllRedeemIncludeSoftDeleted(c echo.Context) error {
	data, err := dc.ds.GetAllIncludeSoftDeleted(c.Request().Context())
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	if len(data) == 0 {
		return response.ResponseSuccess(http.StatusNoContent, data, c)
	} else {
		return response.ResponseSuccess(http.StatusOK, data, c)
	}
}

func (dc redeemController) GetAllRedeemByUserID(c echo.Context) error {
	userID, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := dc.ds.GetAllRedeemByUserID(c.Request().Context(), userID)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	if len(data) == 0 {
		return response.ResponseSuccess(http.StatusNoContent, nil, c)
	} else {
		return response.ResponseSuccess(http.StatusOK, data, c)
	}
}

func (dc redeemController) GetRedeemByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := dc.ds.GetRedeemByID(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "record not found" {
			return response.ResponseSuccess(http.StatusNoContent, nil, c)
		} else {
			return response.ResponseError(http.StatusInternalServerError, err)
		}
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}

func (dc redeemController) CreateRedeem(c echo.Context) error {
	var req dto.RedeemRequest
	c.Bind(&req)

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	userId, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	err = dc.ds.CreateRedeem(c.Request().Context(), req, userId)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusCreated, echo.Map{
		"status": "SUCCESS_INSERT_REDEEM",
	}, c)
}

func (dc redeemController) UpdateRedeem(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	var req dto.RedeemRequest
	c.Bind(&req)

	err = dc.ds.UpdateRedeem(c.Request().Context(), req, id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_UPDATE_PRODUCT",
	}, c)
}

func (dc redeemController) DeleteRedeem(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	err = dc.ds.DeleteRedeem(c.Request().Context(), id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_DELETE_REDEEM",
	}, c)
}
