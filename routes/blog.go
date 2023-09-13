package routes

import (
	"practice/auth/core"

	"github.com/gin-gonic/gin"
)

func SetBlogRoute(router *gin.RouterGroup, application *core.Application) {
	blogModule := application.BlogModule
	blogController := blogModule.Controller

	blogGroup := router.Group("/blog")
	// blogGroup.Use(middlewares.Authentication())
	blogGroup.GET("", blogController.GetBlog)
}
