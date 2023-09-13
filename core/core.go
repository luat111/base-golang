package core

import (
	"practice/auth/core/config"
	"practice/auth/core/database"
	"practice/auth/core/integrates"
	grpcPkg "practice/auth/grpc"
	"practice/auth/modules/blog"
	"practice/auth/modules/post"
	"practice/auth/modules/user"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Application struct {
	*user.UserModule
	*blog.BlogModule
	*post.PostModule
	RabbitMQ    *integrates.Rabbit
	CacheModule *redis.Client
	GRPCMap     map[string]*grpc.ClientConn
}

func initModules(
	databases *database.DatabaseInstance,
	amqp *integrates.Rabbit,
	grpcMap map[string]*grpc.ClientConn,
	serviceMap map[string]string,
) (*Application, error) {

	userModule := user.InitUserModule(databases.SqlDb)
	postModule := post.InitPostModule(databases.MongoDb)
	blogModule := blog.InitBlogModule(grpcMap, serviceMap)

	return &Application{
		CacheModule: databases.CacheDb,
		RabbitMQ:    amqp,
		UserModule:  userModule,
		BlogModule:  blogModule,
		PostModule:  postModule,
		GRPCMap:     grpcMap,
	}, nil
}

func InitApp() (*Application, error) {
	config.InitConfig()

	grpcMap, serviceMap := grpcPkg.InitClientGRPC(config.Config)

	SqlDb := database.ConnectPostgreSQL(config.Config)
	MongoDb := database.ConnectMongoDB(config.Config)
	CacheDb := database.ConnectRedis(config.Config)
	ESClient := database.ConnectElasticSearch(config.Config)
	RMQClient := integrates.ConnectRabbit(config.Config)

	if config.Config.ENV != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}

	return initModules(
		&database.DatabaseInstance{
			SqlDb:    SqlDb,
			MongoDb:  MongoDb,
			CacheDb:  CacheDb,
			ESClient: ESClient,
		},
		RMQClient,
		grpcMap,
		serviceMap,
	)
}
