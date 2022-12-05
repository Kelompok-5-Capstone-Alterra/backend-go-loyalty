package pointController

import (
	pointService "backend-go-loyalty/internal/service/point"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type IPointController interface {
	HandleGetAllPoint(c echo.Context) error
	HandleGetPointByID(c echo.Context) error
}

type pointController struct {
	ps pointService.IPointService
}

func NewPointController(ps pointService.IPointService) pointController {
	return pointController{
		ps: ps,
	}
}

func (pc pointController) HandleGetAllPoint(c echo.Context) error {
	data, err := pc.ps.GetAllPoints(c.Request().Context())
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(data, c)
}

func (pc pointController) HandleGetPointByID(c echo.Context) error {
	userID, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return responseErrorParams(err, c)
	}
	data, err := pc.ps.GetPoint(c.Request().Context(), userID)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(data, c)
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
