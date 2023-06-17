package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/response"
)

type CartHandler struct {
	cartUseCase services.CartUseCase
}

func NewCartHandler(usecase services.CartUseCase) *CartHandler {
	return &CartHandler{
		cartUseCase: usecase,
	}
}

// AddProductToCart
// @Summary User Add Product To Cart By ProductId
// @ID user-add-to-cart-
// @Description User Can add product To Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_id path string true "product_item_id"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/cart/{product_id}  [post]
func (cr *CartHandler) AddToCart(c *gin.Context) {
	pararmsID := c.Param("product_id")
	productID, err := strconv.Atoi(pararmsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	userID := fmt.Sprintf("%v", c.Value("userID"))

	cartItems, err := cr.cartUseCase.AddToCart(c.Request.Context(), userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to add product item to the cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "product added to the cart successfully",
		Data:       cartItems,
		Errors:     nil,
	})
}

// RemoveProductToCart
// @Summary User Remove Product To Cart By ProductId
// @ID user-remove-to-cart-
// @Description User Can Remove product To Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_id path string true "product_item_id"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/remove/{product_id}  [delete]
func (cr *CartHandler) RemoveTOCart(c *gin.Context) {
	paramsID := c.Param("product_id")
	productID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to process request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userID := fmt.Sprintf("%v", c.Value("userID"))
	err = cr.cartUseCase.RemoveTOCart(c.Request.Context(), userID, productID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to remove product from cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "product removed from cart successfuly",
		Data:       nil,
		Errors:     nil,
	})
}

// ViewCarts
// @Summary User Can View Cart and Total Amount
// @ID View-cart
// @Description User Can View Cart and Total Amount
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router  /user/carts  [get]
func (cr *CartHandler) ViewCart(c *gin.Context) {
	userID := fmt.Sprintf("%v", c.Value("userID"))
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, response.Response{
	// 		StatusCode: http.StatusUnauthorized,
	// 		Message:    "unable to fetch userID from context",
	// 		Data:       nil,
	// 		Errors:     err.Error(),
	// 	})
	// 	return
	// }
	cart, err := cr.cartUseCase.ViewCart(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed view from cart details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Your cart is ",
		Data:       cart,
		Errors:     nil,
	})

}

// AddCouponToCart
// @Summary User Can Add coupon To Cart
// @ID add-coupon-to-cart
// @Description User Can Add coupon To Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param couponid path string true "couponid"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 400 {object} response.Response
// @Router  /user/addcoupon/{couponid}  [post]
func (cr *CartHandler) AddCouponToCart(c *gin.Context) {
	userID := fmt.Sprintf("%v", c.Value("userID"))
	paramsID := c.Param("couponid")
	couponID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	cart, err := cr.cartUseCase.AddCouponToCart(c.Request.Context(), userID, couponID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed coupon add to cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "coupon added to cart successfuly",
		Data:       cart,
		Errors:     nil,
	})

}
