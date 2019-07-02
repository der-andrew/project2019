package mongo

import (
	"fmt"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"
)

type Storage struct {
	RawMongo

	storage.Storage
}

type RawMongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func (s *Storage) Raw() interface{} {
	return &s.RawMongo
}

func MustConnect(host string, port int, name, password, database string) *Storage {
	s := &Storage{}

	var client *mongo.Client
	var err error
	if name != "" && password != "" {
		client, err = mongo.NewClient(fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", name, password, host, port, database))
	} else {
		client, err = mongo.NewClient(fmt.Sprintf("mongodb://%s:%d", host, port))
	}

	wrapFatalLog := func(host string, port int) {
		logrus.WithFields(logrus.Fields{
			"context":  "CORE",
			"resource": "STORAGE",
			"addr":     fmt.Sprintf("%s:%d", host, port),
			"status":   "FAILED",
		}).Fatal("DB connection cannot be established")
	}

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"context":     "CORE",
			"description": "Can't create client instance",
			"err":         err,
		}).Debug("Details of the connection error to the database")
		wrapFatalLog(host, port)
	}

	if err = client.Connect(nil); err != nil {
		logrus.WithFields(logrus.Fields{
			"context":     "CORE",
			"description": "Can't connect to mongodb service",
			"err":         err,
		}).Debug("Details of the connection error to the database")
		wrapFatalLog(host, port)
	}

	if err = client.Ping(nil, nil); err != nil {
		logrus.WithFields(logrus.Fields{
			"context":     "CORE",
			"description": "Can't connect to mongodb service",
			"err":         err,
		}).Debug("Details of the connection error to the database")
		wrapFatalLog(host, port)
	}

	s.Client = client
	s.Database = client.Database(database)

	logrus.WithFields(logrus.Fields{
		"resource": "DB",
		"addr":     fmt.Sprintf("%s:%d", host, port),
		"status":   "CONNECTED",
	}).Info("DB connection successfully established")

	return s
}
