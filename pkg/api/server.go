package api

import (
	"github.com/gin-gonic/gin"
	handler "github.com/nadeem1815/premium-watch/pkg/api/handler"
	"github.com/nadeem1815/premium-watch/pkg/api/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler,
) *ServerHTTP {
	engine := gin.New()

	//engine logger from gin
	engine.Use(gin.Logger())

	// swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// request jwt

	userapi := engine.Group("user")

	userapi.POST("/register", userHandler.UserRegister)
	userapi.POST("/login/email", userHandler.LoginWithEmail)

	userapi.Use(middleware.UserAuth)
	userapi.POST("/logout", userHandler.UserLogut)
	userapi.GET("/home", userHandler.Home)

	userapi.GET("/all_product", productHandler.ListAllProducts)

	admin := engine.Group("admin")
	// admin.POST("/register", adminHandler.AdminSingUP)
	admin.POST("/login/email", adminHandler.LoginAdmin)
	// admin.GET()"/logout",)

	admin.Use(middleware.AdminAuth)
	admin.POST("/logout", adminHandler.AdminLogout)
	admin.PATCH("/block_user/:user_id", userHandler.BlockedUser)
	admin.PATCH("/unblock_user/:user_id", userHandler.UnBlockUser)
	admin.GET("/list_all_user", adminHandler.ListAllUsers)
	admin.GET("/find_userid/:user_id", adminHandler.FindUserId)

	admin.POST("/create_categories", productHandler.CreateCategory)
	admin.GET("/all_categories", productHandler.ViewAllCategory)
	admin.GET("/find_category_id/:id", productHandler.FindCategoryById)

	admin.POST("/create_product", productHandler.CreateProduct)
	admin.GET("/all_product", productHandler.ListAllProducts)
	admin.PATCH("/update_product", productHandler.UpdatateProduct)
	admin.DELETE("delete_product/:id", productHandler.DeleteProduct)
	return &ServerHTTP{engine: engine}

}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
