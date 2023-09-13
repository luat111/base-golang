package config

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
)

var (
	Config     AppConfig
	CorsConfig cors.Config
)

func LoadConfig(path string) (config AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	err = viper.Unmarshal(&config.DbSetting)
	err = viper.Unmarshal(&config.MongoSetting)
	err = viper.Unmarshal(&config.CacheSetting)
	err = viper.Unmarshal(&config.GRPCSetting)
	err = viper.Unmarshal(&config.ESSetting)
	err = viper.Unmarshal(&config.RabbitSetting)

	return
}

func GetCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
}

func InitConfig() {
	appConfig, err := LoadConfig("..")
	corsConfig := GetCorsConfig()

	if err != nil {
		panic("Cannot innit app config")
	}

	Config = appConfig
	CorsConfig = corsConfig
}
