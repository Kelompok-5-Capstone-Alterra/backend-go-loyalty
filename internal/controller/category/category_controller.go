package categoryController

import (
	"backend-go-loyalty/internal/dto"
	categoryService "backend-go-loyalty/internal/service/category"
	"backend-go-loyalty/pkg/response"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ICategoryController interface {
	HandleGetAllCategories(c echo.Context) error
	HandleGetCategoryByID(c echo.Context) error
	HandleCreateCategory(c echo.Context) error
	HandleUpdateCategory(c echo.Context) error
	HandleDeleteCategory(c echo.Context) error
}

type categoryController struct {
	cs categoryService.ICategorySerivce
}

func NewCategoryController(cs categoryService.ICategorySerivce) categoryController {
	return categoryController{
		cs: cs,
	}
}

func (cc categoryController) HandleGetAllCategories(c echo.Context) error {
	data, err := cc.cs.GetAllCategories(c.Request().Context())
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (cc categoryController) HandleGetCategoryByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := cc.cs.GetCategoryByID(c.Request().Context(), id)
	if err != nil && err.Error() != "record not found" {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	if err != nil && err.Error() == "record not found" {
		return response.ResponseSuccess(http.StatusOK, nil, c)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (cc categoryController) HandleCreateCategory(c echo.Context) error {
	req := dto.CategoryRequest{}
	c.Bind(&req)
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	err = cc.cs.CreateCategory(c.Request().Context(), req)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCES_INSERT_CATEGORY",
	}, c)
}
func (cc categoryController) HandleUpdateCategory(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	req := dto.CategoryRequest{}
	c.Bind(&req)
	err = cc.cs.UpdateCategory(c.Request().Context(), id, req)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCES_UPDATE_CATEGORY",
	}, c)
}
func (cc categoryController) HandleDeleteCategory(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	err = cc.cs.DeleteCategory(c.Request().Context(), id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCES_DELETE_CATEGORY",
	}, c)
}
