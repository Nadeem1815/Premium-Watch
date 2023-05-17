package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nadeem1815/premium-watch/pkg/domain"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"github.com/nadeem1815/premium-watch/pkg/utils/response"
)

type ProductHandler struct {
	productUseCase services.ProductUseCase
}

func NewProductHandler(productUseCase services.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category model.NewCategory
	if err := c.Bind(&category); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	newCategory, err := cr.productUseCase.CreateCategory(c.Request.Context(), category.CategoryName)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to create new category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: http.StatusCreated,
		Message:    "New category Created",
		Data:       newCategory,
		Errors:     nil,
	})
}

func (cr *ProductHandler) ViewAllCategory(c *gin.Context) {
	categories, err := cr.productUseCase.ViewAllCategory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to fetch all categories",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    " View all Categories",
		Data:       categories,
		Errors:     nil,
	})

}

func (cr *ProductHandler) CreateProduct(c *gin.Context) {
	var createProduct domain.Product
	if err := c.Bind(&createProduct); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "unable to request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	createProducts, err := cr.productUseCase.CreateProduct(c.Request.Context(), createProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to Create Product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Product Created",
		Data:       createProducts,
		Errors:     nil,
	})

}

func (cr *ProductHandler) UpdatateProduct(c *gin.Context) {
	var updataproduct domain.Product
	if err := c.Bind(&updataproduct); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	updateProductItem, err := cr.productUseCase.UpdateProduct(c.Request.Context(), updataproduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to Update Items",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "updated Successfully",
		Data:       updateProductItem,
		Errors:     nil,
	})

}

func (cr *ProductHandler) ListAllProducts(c *gin.Context) {
	listAllProducts, err := cr.productUseCase.ListAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to fetch all products",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "list All Products",
		Data:       listAllProducts,
		Errors:     nil,
	})
}
