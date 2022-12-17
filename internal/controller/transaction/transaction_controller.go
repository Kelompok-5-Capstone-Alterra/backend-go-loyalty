package transactionController

import (
	"backend-go-loyalty/internal/dto"
	transactionService "backend-go-loyalty/internal/service/transaction"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ITransactionController interface {
	HandleGetAllTransaction(c echo.Context) error
	HandleGetTransactionByID(c echo.Context) error
	HandleCreateTransaction(c echo.Context) error
	HandleGetTransactionsByUserID(c echo.Context) error
	HandleGetTransactionByIDByUser(c echo.Context) error
}

type transactionController struct {
	ts transactionService.ITransactionService
}

func NewTransactionController(ts transactionService.ITransactionService) transactionController {
	return transactionController{
		ts: ts,
	}
}

func (tc transactionController) HandleGetAllTransaction(c echo.Context) error {
	data, err := tc.ts.GetAllTransaction(c.Request().Context())
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (tc transactionController) HandleGetTransactionByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := tc.ts.GetTransactionByID(c.Request().Context(), uint64(id))
	if err != nil {
		if err.Error() == "record not found" {
			return response.ResponseSuccess(http.StatusOK, nil, c)
		}
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (tc transactionController) HandleCreateTransaction(c echo.Context) error {
	req := dto.TransactionRequest{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data,err := tc.ts.CreateTransaction(c.Request().Context(), req, id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusCreated, echo.Map{
		"status": "PENDING_PAYMENT",
		"invoice":data,
	}, c)
}
func (tc transactionController) HandleGetTransactionsByUserID(c echo.Context) error {
	id, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := tc.ts.GetTransactionByUserID(c.Request().Context(), id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (tc transactionController) HandleGetTransactionByIDByUser(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	user, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := tc.ts.GetTransactionByIDByUserID(c.Request().Context(), user, uint64(id))
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
