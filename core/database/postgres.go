package database

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"practice/auth/core/config"
	user_model "practice/auth/modules/user/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SqlDB *gorm.DB = nil

func ConnectPostgreSQL(config config.AppConfig) *gorm.DB {
	if SqlDB == nil {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var err error
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.DbSetting.DBHost, config.DbSetting.DBUser, config.DbSetting.DBPass, config.DbSetting.DBName, config.DbSetting.DBPort,
		)

		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the Database! \n", err.Error())
			os.Exit(1)
		}

		connection.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
		connection.Set("gorm:auto_preload", true)

		if config.ENV == "develop" {
			connection.Logger = logger.Default.LogMode(logger.Info)
		}

		log.Println("Running Migrations")
		err = connection.AutoMigrate(
			&user_model.User{},
		)

		if err != nil {
			slog.Info(err.Error())
		}

		log.Println("🚀 Connected Successfully to PostgreSQL Database")
		SqlDB = connection
	}

	return SqlDB
}
