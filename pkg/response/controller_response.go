package response

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ResponseErrorRequestBody(code int, err error) error {
	errRes := ErrorResponseData{}
	for _, val := range err.(validator.ValidationErrors) {
		var errVal ErrorResponseValue
		errVal.Key = val.StructField()
		errVal.Value = val.Tag()
		errRes = append(errRes, errVal)
	}
	return echo.NewHTTPError(code,
		NewBaseResponse(code,
			http.StatusText(code),
			errRes,
			nil))
}

func ResponseError(code int, err error) error {
	var errRes ErrorResponseData
	if code >= 400 {
		errVal := ErrorResponseValue{
			Key:   "error",
			Value: err.Error(),
		}
		errRes = ErrorResponseData{errVal}
	} else {
		errRes = nil
	}
	return echo.NewHTTPError(code,
		NewBaseResponse(code,
			http.StatusText(code),
			errRes,
			nil))
}

func ResponseSuccess(code int, data interface{}, c echo.Context) error {
	return c.JSON(code, NewBaseResponse(code, http.StatusText(code), nil, data))
}
