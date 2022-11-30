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
		return responseErrorInternal(err, c)
	}
	return responseSuccess(data, c)
}

func (rc rewardController) FindRewardById(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorParams(err, c)
	}
	data, err := rc.rs.FindRewardByID(c.Request().Context(), id)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(data, c)
}

func (rc rewardController) CreateReward(c echo.Context) error {
	var req dto.RewardRequest
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return responseErrorValidator(err, c)
	}
	err = rc.rs.CreateReward(c.Request().Context(), req)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(echo.Map{
		"status": "SUCCESS_INSERT_PRODUCT",
	}, c)
}
func (rc rewardController) UpdateReward(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorParams(err, c)
	}
	var req dto.RewardRequest
	c.Bind(&req)
	err = rc.rs.UpdateReward(c.Request().Context(), req, id)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(echo.Map{
		"status": "SUCCESS_UPDATE_PRODUCT",
	}, c)
}

func (rc rewardController) DeleteReward(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorParams(err, c)
	}

	err = rc.rs.DeleteReward(c.Request().Context(), id)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(echo.Map{
		"status": "SUCCESS_DELETE_PRODUCT",
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
