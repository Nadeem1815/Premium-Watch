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
	fmt.Println(userID)
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
