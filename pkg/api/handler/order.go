package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"github.com/nadeem1815/premium-watch/pkg/utils/response"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(orderusecase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: orderusecase,
	}
}

func (cr *OrderHandler) BuyAll(c *gin.Context) {
	var body model.PlaceOrder
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	fmt.Printf("%+v", body)
	UserID := fmt.Sprintf("%v", c.Value("userID"))
	order, err := cr.orderUseCase.BuyAll(c.Request.Context(), body, UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Order failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Order Successfully",
		Data:       order,
		Errors:     nil,
	})

}
