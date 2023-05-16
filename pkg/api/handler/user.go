package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "github.com/nadeem1815/premium-watch/pkg/usecase/interface"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"github.com/nadeem1815/premium-watch/pkg/utils/response"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

// type Response struct {
// 	ID      uint   `copier:"must"`
// 	Name    string `copier:"must"`
// 	Surname string `copier:"must"`
// }

// UserSignUp
func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) Home(c *gin.Context) {
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 200,
		Message:    "Welcome To Home page",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *UserHandler) UserRegister(c *gin.Context) {

	// 1.recive data from request body
	var user model.UsarDataInput
	if err := c.BindJSON(&user); err != nil {
		// Return a 422 Bad request response if the request body is malformed
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "unable to prcess the request",
			Data:       nil,
			Errors:     err.Error(),
		})

		return

	}
	// call the Createuser method of the userUsecase to create the user
	userData, err := cr.userUseCase.UserRegister(c.Request.Context(), user)

	if err != nil {
		// Return a 400 bad request response if there is an error while creating the user.
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to create user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}

	// Return a 201 Created response if the user is successfully created.
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 200,
		Message:    "Created Successfully",
		Data:       userData,
		Errors:     nil,
	})

}

// LoginWithEmail

func (cr *UserHandler) LoginWithEmail(c *gin.Context) {
	// receive data from request body
	var user model.UserLoginEmail
	// return a 421 response if the request body is malformed
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	//  call the userlogin method of the userUseCase to login as a user.
	ss, Users, err := cr.userUseCase.LoginWithEmail(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
	// Return a 200 Success of response if the user is successfully logged in
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Successfully logged in",
		Data:       Users,
		Errors:     nil,
	})

}

func (cr *UserHandler) UserLogut(c *gin.Context) {
	c.Writer.Header().Set("cache-control", "no-cache,no-store,must-revalidate")
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", "", -1, "", "", false, true)
	c.Status(http.StatusOK)

}

func (cr *UserHandler) BlockedUser(c *gin.Context) {
	// var blockedUser model.BlockUser

	// if err := c.Bind(&blockedUser); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, response.Response{
	// 		StatusCode: http.StatusUnprocessableEntity,
	// 		Message:    "Unable to fetch id from context",
	// 		Data:       nil,
	// 		Errors:     err.Error(),
	// 	})
	// 	return
	// }
	paramsId := c.Param("user_id")
	user_id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to find user id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	blockUser, err := cr.userUseCase.BlockUser(c.Request.Context(), user_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to block user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "successfull userblocked",
		Data:       blockUser,
		Errors:     nil,
	})

}

func (cr *UserHandler) UnBlockUser(c *gin.Context) {
	paramsId := c.Param("user_id")
	User_Id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed to parse user id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	UnBlockUser, err := cr.userUseCase.UnBlockUser(c.Request.Context(), User_Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to UnBlock User",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "User UnBlock Successfully",
		Data:       UnBlockUser,
		Errors:     nil,
	})
}
