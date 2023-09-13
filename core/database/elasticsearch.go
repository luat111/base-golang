package database

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"practice/auth/core/config"

	"github.com/elastic/go-elasticsearch/v8"
)

var ESClient *elasticsearch.Client = nil

func ConnectElasticSearch(config config.AppConfig) *elasticsearch.Client {
	if ESClient == nil {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		esUri := fmt.Sprintf(
			"http://%s:%d",
			config.ESSetting.Host,
			config.ESSetting.Port,
		)

		esCnf := elasticsearch.Config{
			Addresses: []string{esUri},
			Username:  config.ESSetting.Username,
			Password:  config.ESSetting.Password,
		}

		client, err := elasticsearch.NewClient(esCnf)

		if err != nil {
			slog.Info(err.Error())
		} else {
			log.Println("🚀 Connected Successfully to ElasticSearch")
			ESClient = client
		}
	}

	return ESClient
}
