package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nadeem1815/premium-watch/pkg/utils/response"
)

func (cr *OrderHandler) UserWallet(c *gin.Context) {
	userID := fmt.Sprintf("%v", c.Value("userID"))

	wallet, err := cr.orderUseCase.UserWallet(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed  fetch user wallet",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "wallet successfully",
		Data:       wallet,
		Errors:     nil,
	})
}
