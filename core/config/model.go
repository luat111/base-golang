package config

type AppConfig struct {
	//Environment
	AppPort                    string `mapstructure:"PORT"`
	ENV                        string `mapstructure:"ENV" json:"ENV"`
	TRUSTED_DOMAIN             string `mapstructure:"TRUSTED_DOMAIN"`
	BASH_PATH                  string `mapstructure:"BASH_PATH"`
	LIMIT_REQUEST              int    `mapstructure:"LIMIT_REQUEST"`
	LIMIT_REQUEST_PER_SECOND   int    `mapstructure:"LIMIT_REQUEST_PER_SECOND"`
	REQUEST_TIMEOUT_PER_SECOND int    `mapstructure:"REQUEST_TIMEOUT_PER_SECOND"`

	//DB CONFIG
	//Postgres
	DbSetting struct {
		DBHost   string `mapstructure:"DB_HOST"`
		DBPort   string `mapstructure:"DB_PORT"`
		DBUser   string `mapstructure:"DB_USER"`
		DBPass   string `mapstructure:"DB_PASS"`
		DBName   string `mapstructure:"DB_NAME"`
		DBSchema string `mapstructure:"DB_SCHEMA"`
	}

	//Mongo
	MongoSetting struct {
		DBHost string `mapstructure:"MONGO_HOST"`
		DBPort string `mapstructure:"MONGO_PORT"`
		DBUser string `mapstructure:"MONGO_USER"`
		DBPass string `mapstructure:"MONGO_PASS"`
		DBName string `mapstructure:"MONGO_DB_NAME"`
	}

	//Redis
	CacheSetting struct {
		CacheHost string `mapstructure:"CACHE_HOST"`
		CachePort string `mapstructure:"CACHE_PORT"`
		CachePass string `mapstructure:"CACHE_PASS"`
		CacheDB   int    `mapstructure:"CACHE_DB"`
	}

	//GRPC
	GRPCSetting struct {
		GRPCPort        string `mapstructure:"GRPC_PORT"`
		BlogServiceHost string `mapstructure:"BLOG_SERVICE_HOST"`
		BlogServicePort string `mapstructure:"BLOG_SERVICE_PORT"`
	}

	//Elastic
	ESSetting struct {
		Host     string `mapstructure:"ES_HOST"`
		Port     int    `mapstructure:"ES_PORT"`
		Username string `mapstructure:"ES_USERNAME"`
		Password string `mapstructure:"ES_PWD"`
	}

	//RMQ
	RabbitSetting struct {
		Host string `mapstructure:"RMQ_HOST"`
		Port string `mapstructure:"RMQ_PORT"`
		User string `mapstructure:"RMQ_USER"`
		Pass string `mapstructure:"RMQ_PWD"`
	}
}
