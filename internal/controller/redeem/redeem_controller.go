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
	GetAllRedeem(c echo.Context) error
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
	data, err := dc.ds.GetAllRedeem(c.Request().Context())
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(data, c)
}

func (dc redeemController) GetRedeemByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorParams(err, c)
	}
	data, err := dc.ds.GetRedeemByID(c.Request().Context(), id)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(data, c)
}

func (dc redeemController) CreateRedeem(c echo.Context) error {
	var req dto.RedeemRequest
	c.Bind(&req)

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return responseErrorValidator(err, c)
	}
	userId, err := utils.GetUserIDFromJWT(c)
	err = dc.ds.CreateRedeem(c.Request().Context(), req, userId)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(echo.Map{
		"status": "SUCCESS_INSERT_REDEEM",
	}, c)
}

func (dc redeemController) UpdateRedeem(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorParams(err, c)
	}
	var req dto.RedeemRequest
	c.Bind(&req)

	err = dc.ds.UpdateRedeem(c.Request().Context(), req, id)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(echo.Map{
		"status": "SUCCESS_UPDATE_PRODUCT",
	}, c)
}

func (dc redeemController) DeleteRedeem(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorParams(err, c)
	}
	err = dc.ds.DeleteRedeem(c.Request().Context(), id)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(echo.Map{
		"status": "SUCCESS_DELETE_REDEEM",
	}, c)
}

func responseErrorInternal(err error, c echo.Context) error {
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

func responseSuccess(result interface{}, c echo.Context) error {
	return c.JSON(http.StatusOK, response.NewBaseResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		result,
	))
}

func responseErrorValidator(err error, c echo.Context) error {
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

func responseErrorParams(err error, c echo.Context) error {
	var errVal response.ErrorResponseValue
	errVal.Key = "error"
	errVal.Value = err.Error()

	return echo.NewHTTPError(http.StatusBadRequest,
		response.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseData(errVal),
			nil))
}
