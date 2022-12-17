package paymentController

import (
	"backend-go-loyalty/internal/dto"
	paymentService "backend-go-loyalty/internal/service/payment"
	transactionService "backend-go-loyalty/internal/service/transaction"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type IPaymentController interface {
	// HandleNotification(c echo.Context) error
	HandlePayWithCredit(c echo.Context) error
	HandlePayWithOVO(c echo.Context) error
}

type paymentController struct {
	ts transactionService.ITransactionService
	ps paymentService.IPaymentService
}

func NewPaymentController(ts transactionService.ITransactionService, ps paymentService.IPaymentService) paymentController {
	return paymentController{
		ts: ts,
		ps: ps,
	}
}

func (pc paymentController) HandlePayWithOVO(c echo.Context) error {
	req := dto.PayWithOVO{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	userID, err := utils.GetUserIDFromJWT(c)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := pc.ps.PayWithOVO(c.Request().Context(), req, userID)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}

func (pc paymentController) HandlePayWithCredit(c echo.Context) error {
	req := dto.PayWithCredit{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusInternalServerError, err)
	}
	userID, err := utils.GetUserIDFromJWT(c)
	err = pc.ps.PayWithCredit(c.Request().Context(), userID, req.TransactionID)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "PAYMENT_SUCCESS",
	}, c)
}

// Webhook
// func (pc paymentController) HandleNotification(c echo.Context) error {
// 	body := make(map[string]interface{}, 0)
// 	c.Bind(&body)
// 	idString := body["external_id"].(string)
// 	transactionID, err := utils.ExtractExternalID(idString)
// 	fmt.Println("id: ", transactionID, " external: ", idString)
// 	if err != nil {
// 		return response.ResponseError(http.StatusBadRequest, err)
// 	}
// 	pretty, _ := json.MarshalIndent(body, "", "  ")
// 	parsedCreatedAt, err := time.Parse(time.RFC3339Nano, body["created"].(string))
// 	if err != nil {
// 		return response.ResponseError(http.StatusBadRequest, err)
// 	}

// 	if body["status"].(string) == "EXPIRED" && parsedCreatedAt.Hour()-time.Now().Hour() >= 1 {
// 		err = pc.ts.UpdateStatus(c.Request().Context(), "EXPIRED", uint64(transactionID))
// 		if err != nil {
// 			return response.ResponseError(http.StatusInternalServerError, err)
// 		}
// 		err = pc.ts.DeleteInvoice(c.Request().Context(), uint64(transactionID))
// 		if err != nil {
// 			return response.ResponseError(http.StatusInternalServerError, err)
// 		}
// 		fmt.Println("======INVOICE EXPIRED======")
// 	} else if body["status"].(string) == "EXPIRED" && parsedCreatedAt.Hour()-time.Now().Hour() < 1 {
// 		fmt.Println("======INVOICE EXPIRED======")
// 	} else {
// 		err = pc.ts.UpdateStatus(c.Request().Context(), "SUCCESS", uint64(transactionID))
// 		if err != nil {
// 			return response.ResponseError(http.StatusInternalServerError, err)
// 		}
// 		fmt.Println("======PAYMENT SUCCESS======")
// 	}
// 	fmt.Println(color.Green(string(pretty)))
// 	fmt.Println("===========================")
// 	return c.JSON(200, echo.Map{
// 		"code": 200,
// 		"data": body,
// 	})
// }
