package controller

import (
	dto "dto/product/product_request"
	echo "labstack/echo/v4"
	"net/http"
	service "service/product_service"
)

type ProductController struct {
	ps service.ProductService
}

func New() ProductController {
	return ProductController{
		ps: &service.ProductService{},
	}
}

func (pc *ProductController) handleGetAll(e echo.Context) error {
	var product []dto.ProductRequest = pc.ps.GetAll()

	return e.JSON(http.StatusOK, product)
}

func (pc *ProductController) handleGetProductByID(e echo.Context) error {
	var id string = e.Param("id")

	product := service.ProductService.GetProductByID(id)

	if product.ID == 0 {
		return e.JSON(http.StatusNotFound, map[string]string{
			"massage": "no product found",
		})
	}

	return e.JSON(http.StatusOK, product)
}

func (pc *ProductController) handleInsertProduct(e echo.Context) error {
	var input *dto.ProductRequest = new(dto.ProductRequest)

	if err := e.Bind(input); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"massage": "invalid request",
		})
	}

	_, product := service.ProductService.InsertProduct(*input)

	return e.JSON(http.StatusCreated, product)
}

func (pc *ProductController) handleUpdateProduct(e echo.Context) error {
	var input *dto.ProductRequest = new(dto.ProductRequest)

	if err := e.Bind(input); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"massage": "invalid request",
		})
	}

	var productId string = e.Param("id")

	_, product := service.ProductService.UpdateProduct(productId, *input)

	return e.JSON(http.StatusOK, product)
}

func (pc *ProductController) handleDeleteProduct(e echo.Context) error {
	var productId string = e.Param("id")

	isSuccess := service.ProductService.DeleteProduct(productId)

	if !isSuccess {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"massage": "failed to delete a product",
		})

	}
	return e.JSON(http.StatusOK, productId)
}
