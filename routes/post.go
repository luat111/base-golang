package routes

import (
	"practice/auth/core"
	. "practice/auth/core/middlewares"
	. "practice/auth/modules/post/model"

	"github.com/gin-gonic/gin"
)

func SetPostRoute(router *gin.RouterGroup, application *core.Application) {
	postModule := application.PostModule
	postController := postModule.Controller

	postGroup := router.Group("/post")
	// postGroup.Use(middlewares.Authentication())
	postGroup.POST("/listing",
		ValidatePaginateRequest[QueryPostSchema, OrderPostSchema](),
		postController.GetListPost,
	)

	postGroup.POST("/create",
		ValidateRequest[CreatePostSchema](),
		postController.CreatePost,
	)

	postGroup.POST("/update/:id",
		ValidateRequest[UpdatePostSchema](),
		postController.UpdatePost,
	)

	postGroup.GET("/:id", postController.GetDetailPost)
	postGroup.DELETE("/:id", postController.DeletePost)
}
