package rewardController

import (
	"backend-go-loyalty/internal/dto"
	rewardService "backend-go-loyalty/internal/service/reward"
	"backend-go-loyalty/pkg/response"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type IRewardController interface {
	FindAllReward(c echo.Context) error
	FindRewardById(c echo.Context) error
	CreateReward(c echo.Context) error
	UpdateReward(c echo.Context) error
	DeleteReward(c echo.Context) error
}

type rewardController struct {
	rs rewardService.IRewardService
}

func NewRewardController(rs rewardService.IRewardService) rewardController {
	return rewardController{
		rs: rs,
	}
}

func (rc rewardController) FindAllReward(c echo.Context) error {
	data, err := rc.rs.FindAllReward(c.Request().Context())
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	if len(data) == 0 {
		return response.ResponseSuccess(http.StatusNoContent, nil, c)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}

func (rc rewardController) FindRewardById(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := rc.rs.FindRewardByID(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "record not found" {
			return response.ResponseSuccess(http.StatusNoContent, nil, c)
		}
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}

func (rc rewardController) CreateReward(c echo.Context) error {
	var req dto.RewardRequest
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	err = rc.rs.CreateReward(c.Request().Context(), req)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusCreated, echo.Map{
		"status": "SUCCESS_INSERT_REWARD",
	}, c)
}
func (rc rewardController) UpdateReward(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	var req dto.RewardRequest
	c.Bind(&req)
	err = rc.rs.UpdateReward(c.Request().Context(), req, id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_UPDATE_REWARD",
	}, c)
}

func (rc rewardController) DeleteReward(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}

	err = rc.rs.DeleteReward(c.Request().Context(), id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_DELETE_REWARD",
	}, c)
}
