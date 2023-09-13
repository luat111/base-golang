package routes

import (
	"practice/auth/core"
	. "practice/auth/core/middlewares"
	. "practice/auth/modules/user/model"

	"github.com/gin-gonic/gin"
)

func SetUserRoute(router *gin.RouterGroup, application *core.Application) {
	userModule := application.UserModule
	userController := userModule.Controller

	userGroup := router.Group("/user")
	// userGroup.Use(middlewares.Authentication())
	userGroup.POST("/listing",
		ValidatePaginateRequest[QueryUserSchema, OrderUserSchema](),
		userController.GetListUser,
	)

	userGroup.POST("/create",
		ValidateRequest[CreateUserSchema](),
		userController.CreateUser,
	)

	userGroup.POST("/update/:id",
		ValidateRequest[UpdateUserSchema](),
		userController.UpdateUser,
	)

	userGroup.GET("/:id", userController.GetDetailUser)
	userGroup.DELETE("/:id", userController.DeleteUser)
}
