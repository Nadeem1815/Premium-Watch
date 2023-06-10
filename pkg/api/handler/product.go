package handler

import (
	"net/http"
	"strconv"

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

func (cr *ProductHandler) FindCategoryById(c *gin.Context) {
	paramID := c.Param("id")
	categoriesid, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "unable to parse category ID",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	categoryId, err := cr.productUseCase.FindCategoryById(c.Request.Context(), categoriesid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "unable to fetch category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Category is",
		Data:       categoryId,
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

func (cr *ProductHandler) DeleteProduct(c *gin.Context) {
	paramsID := c.Param("id")
	id, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed parse product ID",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	err = cr.productUseCase.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to fetch product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Product Deleted",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *ProductHandler) CreateCoupon(c *gin.Context) {
	var couponCreate model.CreatCoupon

	if err := c.Bind(&couponCreate); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to create request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	couponCrt, err := cr.productUseCase.CreateCoupon(c.Request.Context(), couponCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Coupon Creating failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Coupon Created",
		Data:       couponCrt,
		Errors:     nil,
	})
}

func (cr *ProductHandler) UpdateCoupon(c *gin.Context) {
	var updateCoupon model.UpdatCoupon
	if err := c.BindJSON(&updateCoupon); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to create request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	updatedCoupon, err := cr.productUseCase.UpdateCoupon(c.Request.Context(), updateCoupon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "coupon update failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Updated Successfuly",
		Data:       updatedCoupon,
		Errors:     nil,
	})
}

func (cr *ProductHandler) DeleteCoupon(c *gin.Context) {
	parms := c.Param("id")
	couponID, err := strconv.Atoi(parms)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to process request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	err = cr.productUseCase.DeleteCoupon(c.Request.Context(), couponID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed coupon delete",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "coupon deleted Succeffuly",
		Data:       nil,
		Errors:     nil,
	})

}
