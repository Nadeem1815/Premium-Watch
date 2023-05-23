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
