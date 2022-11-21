package pingController

import (
	pingService "backend-go-loyalty/internal/service/ping"
	"backend-go-loyalty/pkg/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingController interface {
	HandlePing(c echo.Context) error
}

type pingController struct {
	ps pingService.PingService
}

func NewPingController(ps pingService.PingService) pingController {
	return pingController{
		ps: ps,
	}
}

func (pc pingController) HandlePing(c echo.Context) error {
	res, _ := pc.ps.GetPing(c.Request().Context())
	return c.JSON(http.StatusOK, response.NewBaseResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		response.NewErrorResponseData(
			response.NewErrorResponseValue("key", "value"),
		),
		res,
	))
}
