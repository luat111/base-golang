package v1

import (
	"net/http"

	"practice/auth/core"
	"practice/auth/core/config"
	. "practice/auth/core/interfaces"
	"practice/auth/core/middlewares"
	"practice/auth/routes"

	"github.com/gin-gonic/gin"
)

func InitRoute(application *core.Application, Config config.AppConfig) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.TimeoutMiddleware(Config.REQUEST_TIMEOUT_PER_SECOND))

	if Config.ENV != "develop" {
		router.Use(
			middlewares.RateLimitMiddleware(
				application.CacheModule,
				Config.LIMIT_REQUEST,
				Config.LIMIT_REQUEST_PER_SECOND,
			))
	}

	basePath := router.Group(Config.BASH_PATH)

	basePath.GET("/ping", func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusOK, ResponseDefault{Status: true, Message: "Pong!"})
	})

	routes.SetUserRoute(basePath, application)
	routes.SetBlogRoute(basePath, application)
	routes.SetPostRoute(basePath, application)

	return router
}
