package pointController

import (
	pointService "backend-go-loyalty/internal/service/point"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IPointController interface {
	HandleGetAllPoint(c echo.Context) error
	HandleGetUserPoint(c echo.Context) error
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
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}

func (pc pointController) HandleGetPointByID(c echo.Context) error {
	id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := pc.ps.GetPoint(c.Request().Context(), id)
	if err != nil && err.Error() != "record not found" {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	if err.Error() == "record not found" {
		return response.ResponseSuccess(http.StatusOK, nil, c)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}

func (pc pointController) HandleGetUserPoint(c echo.Context) error {
	userID, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := pc.ps.GetPoint(c.Request().Context(), userID)
	if err != nil && err.Error() != "record not found" {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	if err.Error() == "record not found" {
		return response.ResponseSuccess(http.StatusOK, nil, c)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)

}
