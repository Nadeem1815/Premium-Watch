package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nadeem1815/premium-watch/pkg/domain"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"github.com/nadeem1815/premium-watch/pkg/utils/response"
)

type AdminHandler struct {
	adminusecase services.AdminUseCase
}

func NewAdminHandler(adminUseCase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminusecase: adminUseCase,
	}
}

// AdminSignUp
// @Summary Admin SignUp
// @ID admin-signup
// @Description New Admin Registration.
// @Tags Admin
// @Accept json
// @Produce json
// @Param newAdmin body domain.Admin true "Register Admin "
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/register [post]
func (cr *AdminHandler) AdminSingUP(c *gin.Context) {
	var newAdmin domain.Admin
	if err := c.Bind(&newAdmin); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to read the request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	// call creatadmin method from admin UseCase
	err := cr.adminusecase.SaveAdmin(c.Request.Context(), newAdmin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to create admin",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}

	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 200,
		Message:    "admin created succefully",
	})
}

// AdminLogin
// @Summary Admin Login
// @ID admin-login
// @Description Admin login
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_credentials body model.AdminLogin true "Admin login credentials"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/login/email  [post]
func (cr *AdminHandler) LoginAdmin(c *gin.Context) {
	// recieve data from request
	var admin model.AdminLogin

	if err := c.Bind(&admin); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed to request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	fmt.Println(admin)
	// call the Adminlogin method of adminUseCase login as admin
	ss, admins, err := cr.adminusecase.AdminLogin(c.Request.Context(), admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuth", ss, 36000*24*30, "", "", false, true)
	// return 200 a Success response if the admin will logged Successfully
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 200,
		Message:    "Successfully logged",
		Data:       admins,
		Errors:     nil,
	})

}

// AdminLogout
// @Summary Admin Logout
// @ID admin-logout
// @Description Logs out a logged-in admin from the E-commerce web api admin panel
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /admin/logout [post]
func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.Writer.Header().Set("cache-control", "no-cache,no-store,must-revalidate")
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuth", "", -1, "", "", false, true)
	c.Status(http.StatusOK)
}

// func (cr *AdminHandler) BlockUser(c *gin.Context) {
// 	var blockUser model.BlockUser
// 	if err := c.Bind(&blockUser); err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "Unable to fetch admin request to blockuser",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}
// }

// ListAllUsers
// @Summary List All Users
// @ID list-all-users
// @Description Admin Can List All Register Users
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router  /admin/list_all_user  [get]
func (cr *AdminHandler) ListAllUsers(c *gin.Context) {
	listUser, err := cr.adminusecase.ListAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "unable to fetch all Users",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "list of all Users",
		Data:       listUser,
		Errors:     nil,
	})
}

// FindUserById
// @Summary Find User By Id
// @ID find-user-id
// @Description Admin Can Find All Register Users details find By user id
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be fetched"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/find_userid/:user_id  [get]
func (cr *AdminHandler) FindUserId(c *gin.Context) {
	paramsId := c.Query("user_id")
	// fmt.Println(paramsId, "///////////")
	// userId, err := strconv.Atoi(paramsId)
	// fmt.Println(userId, "////////")
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, response.Response{
	// 		StatusCode: http.StatusUnprocessableEntity,
	// 		Message:    "failed parse user id",
	// 		Data:       nil,
	// 		Errors:     err.Error(),
	// 	})
	// 	return
	// }
	user, err := cr.adminusecase.FindUserId(c.Request.Context(), paramsId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "filed  fetch userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "user Id",
		Data:       user,
		Errors:     nil,
	})
}

// DashBoard
// @Summary DashBoard
// @ID dash_board
// @Description Admin Can access dashboard and view details of recoding orders,products etc
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router  /admin/dashboard  [get]
func (cr *AdminHandler) DashBoard(c *gin.Context) {
	dashBoard, err := cr.adminusecase.DashBoard(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed fetch admin dashboard",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "AdminDashBoard",
		Data:       dashBoard,
		Errors:     nil,
	})
}

// SalesReport
// @Summary Sales Report
// @ID sales_report
// @Description Admin can download total sales report in csv.format
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router  /admin/download  [get]
func (cr *AdminHandler) SalesRepo(c *gin.Context) {
	salesReport, err := cr.adminusecase.SalesRepo(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to read request body ",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	// set headers for downloading in browser
	// Set the appropriate headers for the download
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=premiumwatch.csv")

	// Create a CSV writer using our response writer as our io.Writer
	wr := csv.NewWriter(c.Writer)

	// Write CSV header row
	headers := []string{"OrderID", "UserID", "Total", "CouponCode", "Payment Method", "Order Status", "Delivery Status", "Order Date"}
	if err := wr.Write(headers); err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to generate sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	if err := wr.Error(); err != nil {

		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to generate sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	for _, sales := range salesReport {
		row := []string{
			fmt.Sprintf("%v", sales.OrderID),
			sales.UserID,
			fmt.Sprintf("%v", sales.Total),
			sales.CouponCode,
			sales.PaymentMethod,
			sales.OrderStatus,
			sales.DeliveryStatus,
			sales.OrderDate.Format("2006-01-02 15:04:05")}
		fmt.Println("delivery:", sales.DeliveryStatus)
		if err := wr.Write(row); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	// Flush the writer's buffer to ensure all data is written to the client
	wr.Flush()
}
