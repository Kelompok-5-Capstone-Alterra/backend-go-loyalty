package productController

import (
	"backend-go-loyalty/internal/dto"
	productService "backend-go-loyalty/internal/service/product"
	"backend-go-loyalty/pkg/response"
	"strconv"

	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type IProductController interface {
	GetAll(c echo.Context) error
	InsertProduct(c echo.Context) error
	GetProductById(c echo.Context) error
	UpdateProduct(c echo.Context) error
	DeleteProduct(c echo.Context) error
}

type productController struct {
	ps productService.IProductService
}

func NewProductController(ps productService.IProductService) productController {
	return productController{
		ps: ps,
	}
}

func (pc productController) GetAll(c echo.Context) error {
	data, err := pc.ps.GetAll(c.Request().Context())
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}

	var code int
	if len(data) == 0 {
		code = http.StatusNoContent
	} else {
		code = http.StatusOK
	}
	return response.ResponseSuccess(code, data, c)
}
func (pc productController) InsertProduct(c echo.Context) error {
	var req dto.ProductRequest
	c.Bind(&req)

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return response.ResponseErrorRequestBody(http.StatusBadRequest, err)
	}
	err = pc.ps.InsertProduct(c.Request().Context(), req)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusCreated, echo.Map{"status": "SUCCESS_INSERT_PRODUCT"}, c)
}
func (pc productController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	data, err := pc.ps.GetProductByID(c.Request().Context(), id)
	if err != nil {
		if err.Error() == "record not found" {
			return response.ResponseSuccess(http.StatusNoContent, nil, c)
		}
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, data, c)
}
func (pc productController) UpdateProduct(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	var req dto.ProductUpdateRequest
	c.Bind(&req)

	err = pc.ps.UpdateProduct(c.Request().Context(), req, id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_UPDATE_PRODUCT",
	}, c)
}
func (pc productController) DeleteProduct(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return response.ResponseError(http.StatusBadRequest, err)
	}
	err = pc.ps.DeleteProduct(c.Request().Context(), id)
	if err != nil {
		return response.ResponseError(http.StatusInternalServerError, err)
	}
	return response.ResponseSuccess(http.StatusOK, echo.Map{
		"status": "SUCCESS_DELETE_PRODUCT",
	}, c)
}
