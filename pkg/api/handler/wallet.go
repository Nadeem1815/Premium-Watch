package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/nadeem1815/premium-watch/pkg/utils/model"
	"github.com/nadeem1815/premium-watch/pkg/utils/response"
)

// UserWallet
// @Summary User Wallet
// @ID user-wallet
// @Description user
// @Tags User
// @Accept json
// @Produce json
// @Param admin_credentials body model.AdminLogin true "Admin login credentials"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/login/email  [post]
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
