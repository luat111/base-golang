package database

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type DatabaseInstance struct {
	SqlDb    *gorm.DB
	MongoDb  *mongo.Database
	CacheDb  *redis.Client
	ESClient *elasticsearch.Client
}
