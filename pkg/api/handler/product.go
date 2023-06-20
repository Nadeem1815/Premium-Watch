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

// Category management
// CreateCategory
// @Summary Create new product category
// @ID product-category
// @Description Admin can create new category from admin panel
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_name body model.NewCategory true "New category name"
// @Success 201 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router    /admin/create_categories [post]
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
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
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

// ViewAllCategory
// @Summary View All category
// @ID view-all-category
// @Description Admin, users and unregistered users can see all the available categories
// @Tags Product Category
// @Accept json
// @Produce json
// @Success 201 {object} response.Response
// @Failure 500 {object} response.Response
// @Router  /admin/all_categories [get]
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

// FindCategoryByID
// @Summary Find Category by id
// @ID find-category-id
// @Description Admin, users and unregistered users can see all the available categories
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_id path int true "find category by id"
// @Success 201 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 422 {object} response.Response
// @Router  /admin/find_category_id/{id} [get]
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

// Create Product
// @Summary Create new product
// @ID create-product
// @Description Admin can create new products listing
// @Tags Products
// @Accept json
// @Produce json
// @Param createproduct_details body domain.Product true "New product name"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router    /admin/create_product [post]
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

// Update Product
// @Summary Update product
// @ID update-product
// @Description Admin Update products details
// @Tags Products
// @Accept json
// @Produce json
// @Param updateproduct_details body domain.Product true "update product "
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router  /admin/update_product [patch]
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

// ListAllProducts
// @Summary List All Products
// @ID list-all-products
// @Description Admin, users and unregistered users can see all the available products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router  /admin/all_product [get]
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

// DeleteProduct
// @Summary Admin Remove Product To Cart
// @ID Delete-product
// @Description This endpoint allows an admin user to delete a product by ID.
// @Tags Products
// @Accept json
// @Produce json
// @Param product_id path int true "product_id"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/delete_product/{id}  [delete]
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

// Create Coupon
// @Summary Create new Coupon
// @ID create-coupon
// @Description Admin can create new coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param createcoupon_details body model.CreatCoupon true "New Coupon"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router  /admin/creatcoupon [post]
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

// Update Coupon
// @Summary Update Coupon
// @ID update-coupon
// @Description Admin Update Coupon details
// @Tags Coupon
// @Accept json
// @Produce json
// @Param updatecoupon_details body model.UpdatCoupon true "update product "
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router  /admin/updatecoupon [patch]
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

// DeleteCoupon
// @Summary Admin Remove Coupon
// @ID Delete-Coupon
// @Description This endpoint allows an admin user to delete a product by ID.
// @Tags Coupon
// @Accept json
// @Produce json
// @Param coupon_id path int true "coupon_id"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/delete/{id}  [delete]
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

// ViewAllCoupon
// @Summary List All Coupon
// @ID view-all-coupon
// @Description Admin, users and unregistered users can see all the available Coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router  /admin/view [get]
func (cr *ProductHandler) ViewAllCoupon(c *gin.Context) {
	ViewAllCoupon, err := cr.productUseCase.ViewAllCoupon()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to fetch ViewAllCoupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "All coupons",
		Data:       ViewAllCoupon,
		Errors:     nil,
	})

}

// ViewCouponByID
// @Summary View Coupon by id
// @ID view-couponby-id
// @Description Admin, users and registered users can see all the available coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param viewcoupon_id path int true "find coupon by id"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router  /user/coupon/{couponid} [get]
func (cr *ProductHandler) ViewCouponById(c *gin.Context) {
	paramID := c.Param("couponid")
	couponID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed parse couponId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	couponInfo, err := cr.productUseCase.ViewCouponById(c.Request.Context(), couponID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed fetch coupon ",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "coupon is ",
		Data:       couponInfo,
		Errors:     nil,
	})
}
