package integration

import (
	"sync"

	"github.com/spf13/viper"

	"gitlab-host/maximus-platform/thesaurus/service"
	"gitlab-host/maximus-platform/thesaurus/storage"
	"gitlab-host/maximus-platform/thesaurus/storage/mongo"
)

var Storage storage.Storage
var Service *service.Service

var storageOnce, serviceOnce sync.Once

func InitService() {
	InitStorage()

	serviceOnce.Do(func() {
		Service = service.NewInstance(Storage)
	})
}

func InitStorage() {
	storageOnce.Do(func() {
		Storage = mongo.MustConnect(
			viper.GetString("db.host"),
			viper.GetInt("db.port"),
			viper.GetString("db.login"),
			viper.GetString("db.password"),
			"TESTED",
		)
	})
}
