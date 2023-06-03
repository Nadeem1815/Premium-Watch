package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"github.com/nadeem1815/premium-watch/pkg/utils/response"
)

type PaymentHandler struct {
	paymentUseCase services.PaymentUseCase
}

func NewPaymentHandler(payment services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: payment,
	}
}

func (cr *PaymentHandler) CreateRazorPayment(c *gin.Context) {
	paramsId := c.Param("order_id")
	orderID, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to read orderId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	userID := fmt.Sprintf("%v", c.Value("userID"))
	fmt.Println(userID)
	order, razorPayID, err := cr.paymentUseCase.CreateRazorPayment(c.Request.Context(), userID, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "order failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.HTML(200, "app.html", gin.H{
		"UserID":      order.UserID,
		"total_price": order.OrderTotal,
		"total":       order.OrderTotal,
		"orderData":   order.ID,
		"orderid":     razorPayID,
		//"orderid":      order.ID,
		"amount":       order.OrderTotal,
		"Email":        "nadeemf408@gmail.com",
		"Phone_Number": "8129487958",
	})

}

func (cr *PaymentHandler) PaymentSuccess(c *gin.Context) {
	paymentRef := c.Query("payment_ref")
	idstr := c.Query("order_id")
	idstr = strings.ReplaceAll(idstr, "", "")
	orderID, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "can't find orderID",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	uID := c.Query("user_id")
	t := c.Query("total")
	total, err := strconv.ParseFloat(t, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to fetch total from request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	paymentVarifier := model.PaymentVarification{
		UserID:     uID,
		OrderID:    orderID,
		PaymentRef: paymentRef,
		Total:      total,
	}
	err = cr.paymentUseCase.UpatePaymentDetails(c.Request.Context(), paymentVarifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed payment details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "payment Success",
		Data:       true,
		Errors:     nil,
	})
}
