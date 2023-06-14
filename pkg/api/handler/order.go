package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

func (cr *OrderHandler) UserCancelOrder(c *gin.Context) {
	paramsId := c.Param("oderid")
	orderID, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to parse orderID",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	fmt.Println("-----", orderID)

	UserID := fmt.Sprintf("%v", c.Value("userID"))
	cancelOrder, err := cr.orderUseCase.CancelOrder(c.Request.Context(), orderID, UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed order cancel",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "order cancel successfuly",
		Data:       cancelOrder,
		Errors:     nil,
	})

}

func (cr *OrderHandler) UpdateOrder(c *gin.Context) {

	var body model.UpdateOrder
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "cant bind read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.orderUseCase.UpdateOrder(c.Request.Context(), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Update failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "updated Successfuly",
		Data:       order,
		Errors:     nil,
	})

}

func (cr *OrderHandler) ViewAllOrder(c *gin.Context) {
	UserID := fmt.Sprintf("%v", c.Value("userID"))
	viewOrder, err := cr.orderUseCase.ViewAllOrder(c.Request.Context(), UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "faild fetch to order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Your Order is",
		Data:       viewOrder,
		Errors:     nil,
	})
}

func (cr *OrderHandler) ViewOrderID(c *gin.Context) {
	paramsID := c.Param("order_id")
	orderID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "unable to parse orderId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userID := fmt.Sprintf("%v", c.Value("userID"))
	order, err := cr.orderUseCase.ViewOrderID(c.Request.Context(), userID, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to fetch ordersByid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "your order is",
		Data:       order,
		Errors:     nil,
	})
}

func (cr *OrderHandler) RetrunReq(c *gin.Context) {
	var orderId model.RetrunRequest
	if err := c.Bind(&orderId); err != nil {

		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userID := fmt.Sprintf("%v", c.Value("userID"))
	returnRequest, err := cr.orderUseCase.ReturnReq(c.Request.Context(), userID,orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "return failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Return Request Successfuly",
		Data:       returnRequest,
		Errors:     nil,
	})

}
