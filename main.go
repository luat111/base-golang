package main

import (
	"practice/auth/core"
	"practice/auth/core/config"
	. "practice/auth/core/constants/amqp"
	"practice/auth/docs"
	v1 "practice/auth/routes/v1"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	application, _ := core.InitApp()
	defer shutdownApp(application)

	initRoutingAMQP(application)

	appCnf := config.Config
	corsCnf := config.CorsConfig

	r := v1.InitRoute(application, appCnf)
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{appCnf.TRUSTED_DOMAIN})
	r.Use(cors.New(corsCnf))

	gateway := appCnf.TRUSTED_DOMAIN + ":" + appCnf.AppPort
	docs.SwaggerInfo.Host = gateway
	docs.SwaggerInfo.BasePath = "/" + appCnf.BASH_PATH

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(gateway)
}

func shutdownApp(application *core.Application) {
	application.RabbitMQ.Connection.Close()
	application.CacheModule.Close()
}

func initRoutingAMQP(application *core.Application) {
	exchanges := []string{Exchange1, Exchange2}
	queues := []string{Queue1, Queue2}

	application.RabbitMQ.InitExchange(exchanges)
	application.RabbitMQ.InitQueue(queues)
	application.RabbitMQ.BindingQueuesToExchange(queues, Exchange1, "")

	application.RabbitMQ.StartConsume(Queue1)
	application.RabbitMQ.StartConsume(Queue2)
}
