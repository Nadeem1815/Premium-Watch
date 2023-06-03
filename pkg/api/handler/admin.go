package handler

import (
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
