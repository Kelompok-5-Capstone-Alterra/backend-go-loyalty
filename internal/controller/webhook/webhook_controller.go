package webhookController

import (
	transactionService "backend-go-loyalty/internal/service/transaction"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IWebhookController interface {
	HandleEwalletPaymentCallback(c echo.Context) error
}

type webhookController struct {
	ts transactionService.ITransactionService
}

func NewWebhookController(ts transactionService.ITransactionService) webhookController {
	return webhookController{
		ts: ts,
	}
}

func (wc webhookController) HandleEwalletPaymentCallback(c echo.Context) error {
	payload := make(map[string]interface{})
	c.Bind(&payload)
	fmt.Println(payload)
	data := payload["data"].(map[string]interface{})
	transactionID, err := utils.ExtractExternalID(data["reference_id"].(string))
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	err = wc.ts.UpdateStatus(c.Request().Context(), data["status"].(string), uint64(transactionID))
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, nil, c)
}
