package faqController

import (
	"backend-go-loyalty/internal/dto"
	faqService "backend-go-loyalty/internal/service/faq"
	"backend-go-loyalty/pkg/response"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type IFaqController interface {
	HandleGetAllFAQByKeyword(c echo.Context) error
	HandleGetFAQByID(c echo.Context) error
	HandleCreateFAQ(c echo.Context) error
	HandleUpdateFAQ(c echo.Context) error
	HandleDeleteFAQ(c echo.Context) error
}

type faqController struct {
	fs faqService.IFaqService
}

func NewFAQController(fs faqService.IFaqService) faqController {
	return faqController{
		fs: fs,
	}
}

func (fc faqController) HandleGetAllFAQByKeyword(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	data, err := fc.fs.GetAllFAQ(c.Request().Context(), keyword)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (fc faqController) HandleGetFAQByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := fc.fs.GetFAQByID(c.Request().Context(), id)
	if err != nil && err.Error() != "record not found" {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	if err != nil && err.Error() == "record not found" {
		return response.ResponseSuccess(http.StatusOK, nil, c)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (fc faqController) HandleCreateFAQ(c echo.Context) error {
	req := dto.FAQRequest{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	err = fc.fs.CreateFAQ(c.Request().Context(), req)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusCreated, echo.Map{
		"status": "SUCCESS_CREATE_FAQ",
	}, c)
}
func (fc faqController) HandleUpdateFAQ(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	req := dto.FAQUpdateRequest{}
	c.Bind(&req)
	err = fc.fs.UpdateFAQ(c.Request().Context(), req, id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_UPDATE_FAQ",
	}, c)
}
func (fc faqController) HandleDeleteFAQ(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	err = fc.fs.DeleteFAQ(c.Request().Context(), id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_DELETE_FAQ",
	}, c)
}
