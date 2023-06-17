package handler

import (
	"fmt"
	"net/http"

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

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// HomePage
// @Summary HomePage
// @ID home
// @Description when login user can see home page
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/home  [get]
func (cr *UserHandler) Home(c *gin.Context) {
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 200,
		Message:    "Welcome To Home page",
		Data:       nil,
		Errors:     nil,
	})
}

// UserSignUp
// @Summary User SignUp
// @ID user-signup
// @Description New User  can Registration.
// @Tags User
// @Accept json
// @Produce json
// @Param user_details body model.UsarDataInput true "Register User"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /user/register [post]
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
// @Summary LoginWithEmail
// @ID login-with-email
// @Description User login With Email
// @Tags User
// @Accept json
// @Produce json
// @Param User_details body model.UserLoginEmail true "User login credentials"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /user/login/email  [post]
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

// UserLogout
// @Summary User Logout
// @ID user-logout
// @Description Logs out a logged-in user from the E-commerce web api user panel
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /user/logout [post]
func (cr *UserHandler) UserLogut(c *gin.Context) {
	c.Writer.Header().Set("cache-control", "no-cache,no-store,must-revalidate")
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", "", -1, "", "", false, true)
	c.Status(http.StatusOK)

}

// AdminBlockedUser
// @Summary Admin Blocked USer
// @ID admin-block-user
// @Description Admin Blocked for user
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be blocked"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/block_user/:user_id  [patch]
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
	// user_id, err := strconv.Atoi(paramsId)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, response.Response{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "failed to find user id",
	// 		Data:       nil,
	// 		Errors:     err.Error(),
	// 	})
	// 	return
	// }

	blockUser, err := cr.userUseCase.BlockUser(c.Request.Context(), paramsId)
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

// AdminUnBlockedUser
// @Summary Admin UnBlocked USer
// @ID admin-unblock-user
// @Description Admin UnBlocked for user
// @Tags User
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be Unblocked"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /admin/unblock_user/:user_id  [patch]
func (cr *UserHandler) UnBlockUser(c *gin.Context) {
	paramsId := c.Param("user_id")
	// User_Id, err := strconv.Atoi(paramsId)
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, response.Response{
	// 		StatusCode: http.StatusUnprocessableEntity,
	// 		Message:    "failed to parse user id",
	// 		Data:       nil,
	// 		Errors:     err.Error(),
	// 	})
	// 	return

	// }
	UnBlockUser, err := cr.userUseCase.UnBlockUser(c.Request.Context(), paramsId)
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

// AddAddressUser
// @Summary User Add Address
// @ID user-addaddress
// @Description User add address field
// @Tags User
// @Accept json
// @Produce json
// @Param user_address body model.AddressInput true "User add Address"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 422 {object} response.Response
// @Router  /user/address  [post]
func (cr *UserHandler) AddAddress(c *gin.Context) {
	var body model.AddressInput

	if err := c.Bind(&body); err != nil {

		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Unable read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userID := fmt.Sprintf("%v", c.Value("userID"))
	address, err := cr.userUseCase.AddAddress(c.Request.Context(), body, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to add address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: http.StatusOK,
		Message:    "Address created successfuly",
		Data:       address,
		Errors:     nil,
	})

}

