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
		return responseErrorInternal(err, c)
	}

	return responseSuccess(data, c)
}
func (pc productController) InsertProduct(c echo.Context) error {
	var req dto.ProductRequest
	c.Bind(&req)

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return responseErrorValidator(err, c)
	}
	err = pc.ps.InsertProduct(c.Request().Context(), req)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(echo.Map{
		"status": "SUCCESS_INSERT_PRODUCT",
	}, c)
}
func (pc productController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorParams(err, c)
	}
	data, err := pc.ps.GetProductByID(c.Request().Context(), id)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(data, c)
}
func (pc productController) UpdateProduct(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorParams(err, c)
	}
	var req dto.ProductUpdateRequest
	c.Bind(&req)

	err = pc.ps.UpdateProduct(c.Request().Context(), req, id)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(echo.Map{
		"status": "SUCCESS_UPDATE_PRODUCT",
	}, c)
}
func (pc productController) DeleteProduct(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return responseErrorParams(err, c)
	}
	err = pc.ps.DeleteProduct(c.Request().Context(), id)
	if err != nil {
		return responseErrorInternal(err, c)
	}
	return responseSuccess(echo.Map{
		"status": "SUCCESS_DELETE_PRODUCT",
	}, c)
}

func responseErrorInternal(err error, c echo.Context) error {
	errVal := response.ErrorResponseValue{
		Key:   "error",
		Value: err.Error(),
	}
	errRes := response.ErrorResponseData{errVal}
	return echo.NewHTTPError(http.StatusInternalServerError,
		response.NewBaseResponse(http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			errRes,
			nil))
}

func responseSuccess(result interface{}, c echo.Context) error {
	return c.JSON(http.StatusOK, response.NewBaseResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		nil,
		result,
	))
}

func responseErrorValidator(err error, c echo.Context) error {
	errRes := response.ErrorResponseData{}
	for _, val := range err.(validator.ValidationErrors) {
		var errVal response.ErrorResponseValue
		errVal.Key = val.StructField()
		errVal.Value = val.Tag()
		errRes = append(errRes, errVal)
	}
	return echo.NewHTTPError(http.StatusBadRequest,
		response.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			errRes,
			nil))
}

func responseErrorParams(err error, c echo.Context) error {
	var errVal response.ErrorResponseValue
	errVal.Key = "error"
	errVal.Value = err.Error()

	return echo.NewHTTPError(http.StatusBadRequest,
		response.NewBaseResponse(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
			response.NewErrorResponseData(errVal),
			nil))
}

// func InsertProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var product entity.Product
// 	json.NewDecoder(r.Body).Decode(&product)
// 	Instance.Create(&product)
// 	json.NewEncoder(w).Encode(product)
// }

// func GetAll(w http.ResponseWriter, r *http.Request) {
// 	var products []entity.Product
// 	Instance.Find(&products)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(products)
// }

// func GetProductById(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if !checkIfProductExists(productId) {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entity.Product
// 	Instance.First(&product, productId)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	productId := mux.Vars(r)["id"]
// 	if !checkIfProductExists(productId) {
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entity.Product
// 	Instance.First(&product, productId)
// 	json.NewDecoder(r.Body).Decode(&product)
// 	Instance.Save(&product)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(product)
// }

// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	productId := mux.Vars(r)["id"]
// 	if !checkIfProductExists(productId) {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Product Not Found!")
// 		return
// 	}
// 	var product entity.Product
// 	Instance.Delete(&product, productId)
// 	json.NewEncoder(w).Encode("Product Deleted Successfully!")
// }

// func checkIfProductExists(productId string) bool {
// 	var product entity.Product
// 	Instance.First(&product, productId)
// 	if product.ProductID == 0 {
// 		return false
// 	}
// 	return true
// }
