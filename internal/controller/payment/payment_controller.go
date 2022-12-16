package paymentController

import (
	transactionService "backend-go-loyalty/internal/service/transaction"
	"backend-go-loyalty/pkg/response"
	"backend-go-loyalty/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
)

type IPaymentController interface {
	HandleNotification(c echo.Context) error
}

type paymentController struct {
	ts transactionService.ITransactionService
}

func NewPaymentController(ts transactionService.ITransactionService) paymentController {
	return paymentController{
		ts: ts,
	}
}

func (pc paymentController) HandleNotification(c echo.Context) error {
	body := make(map[string]interface{}, 0)
	c.Bind(&body)
	idString := body["external_id"].(string)
	transactionID, err := utils.ExtractExternalID(idString)
	fmt.Println("id: ", transactionID, " external: ", idString)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	pretty, _ := json.MarshalIndent(body, "", "  ")
	parsedCreatedAt, err := time.Parse(time.RFC3339Nano, body["created"].(string))
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}

	if body["status"].(string) == "EXPIRED" && parsedCreatedAt.Hour()-time.Now().Hour() >= 1 {
		err = pc.ts.UpdateStatus(c.Request().Context(), "EXPIRED", uint64(transactionID))
		if err != nil {
			return response.ResponseError(http.StatusInternalServerError, err)
		}
		err = pc.ts.DeleteInvoice(c.Request().Context(), uint64(transactionID))
		if err != nil {
			return response.ResponseError(http.StatusInternalServerError, err)
		}
		fmt.Println("======INVOICE EXPIRED======")
	} else if body["status"].(string) == "EXPIRED" && parsedCreatedAt.Hour()-time.Now().Hour() < 1 {
		fmt.Println("======INVOICE EXPIRED======")
	} else {
		err = pc.ts.UpdateStatus(c.Request().Context(), "SUCCESS", uint64(transactionID))
		if err != nil {
			return response.ResponseError(http.StatusInternalServerError, err)
		}
		fmt.Println("======PAYMENT SUCCESS======")
	}
	fmt.Println(color.Green(string(pretty)))
	fmt.Println("===========================")
	return c.JSON(200, echo.Map{
		"code": 200,
		"data": body,
	})
}
