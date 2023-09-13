package database

import (
	"context"
	"log"
	"practice/auth/core/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var CacheDB *redis.Client = nil

func ConnectRedis(config config.AppConfig) *redis.Client {
	contextBackGround := context.Background()

	if CacheDB == nil {

		_, cancel := context.WithTimeout(contextBackGround, 10*time.Second)
		defer cancel()

		client := redis.NewClient(&redis.Options{
			Addr:     config.CacheSetting.CacheHost + ":" + config.CacheSetting.CachePort,
			Password: config.CacheSetting.CachePass,
			DB:       config.CacheSetting.CacheDB,
		})

		_, err := client.Ping(contextBackGround).Result()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("🗃️  Connected Successfully to the Redis")
		CacheDB = client
	}
	return CacheDB
}
