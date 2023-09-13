package database

import (
	"context"
	"fmt"
	"log"
	"practice/auth/core/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database = nil

func ConnectMongoDB(config config.AppConfig) *mongo.Database {
	if MongoDB == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		connectString := fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/?directConnection=true",
			config.MongoSetting.DBUser, config.MongoSetting.DBPass, config.MongoSetting.DBHost, config.MongoSetting.DBPort,
		)

		dbName := config.MongoSetting.DBName

		clientOptions := options.Client().ApplyURI(connectString)
		client, err := mongo.Connect(ctx, clientOptions)

		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("🚀 Connected Successfully to MongoDB Database")
			MongoDB = client.Database(dbName)
			return MongoDB
		}
	}
	return MongoDB
}
