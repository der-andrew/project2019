package handlers

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	log "github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/models"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/restapi/operations/status"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"
	mongoStorage "repo.nefrosovet.ru/maximus-platform/thesaurus/storage/mongo"
)

func StatusView(_ status.StatusViewParams) middleware.Responder {
	responseInternalServerError := func(logMessage string) middleware.Responder {
		log.WithFields(log.Fields{
			"context": "API",
			"error":   logMessage,
		}).Error(InternalServerErrorMessage)

		payload := new(status.StatusViewInternalServerErrorBody)
		payload.Version = &Version
		payload.Message = &InternalServerErrorMessage

		return status.NewStatusViewInternalServerError().WithPayload(payload)
	}

	responseSuccess := func(statusObjects []models.StatusObject) middleware.Responder {
		payload := new(status.StatusViewOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage

		for _, statusObject := range statusObjects {
			payload.Data = append(payload.Data, &status.DataItems0{
				StatusObject: statusObject,
			})
		}

		return status.NewStatusViewOK().WithPayload(payload)
	}

	collections, err := getCollections()
	if err != nil {
		return responseInternalServerError(err.Error())
	}

	statusObjects := []models.StatusObject{}
	for _, val := range collections {
		collection := Service.Raw().(*mongoStorage.RawMongo).Database.Collection(val)

		sLocales, err := asStrings(
			extractField(collection, &bson.M{"locale": 1, "_id": 0}),
		)
		if err != nil {
			return responseInternalServerError(err.Error())
		}

		locales := []string{}
		// Remove duplicates
		for _, locale := range sLocales {
			if !contain(locales, locale) {
				locales = append(locales, locale)
			}
		}

		if len(locales) != 0 {
			metadata, err := Service.GetMetadata(collection.Name())
			if err != nil {
				switch err {
				case storage.ErrMetadataNotFound:
					log.WithFields(log.Fields{
						"context": "API",
						"error":   err.Error(),
					})
					continue
				default:
					return responseInternalServerError(err.Error())
				}
			}

			statusObjects = append(statusObjects, models.StatusObject{
				Name:    val,
				Locales: locales,
				Type:    string(metadata.Type),
				Updated: metadata.DateOfChange,
			})
		}
	}

	return responseSuccess(statusObjects)
}

func getCollections() ([]string, error) {
	cursor, err := Service.Raw().(*mongoStorage.RawMongo).Database.ListCollections(nil, bson.M{})
	if err != nil {
		log.WithError(err).Error("Error while getting cursor")
		return nil, err
	}

	var collections []string

	collection := bson.M{}
	for cursor.Next(nil) {
		if err := cursor.Decode(&collection); err != nil {
			log.WithError(err).Error("Error while cursor iterate")
			return nil, err
		}

		if collection["name"] == "_metadata" { // skip metadata
			continue
		}

		collections = append(collections, collection["name"].(string))
	}

	return collections, nil
}

func extractField(collection *mongo.Collection, projection *bson.M) []interface{} {
	cursor, err := collection.Find(nil, bson.M{
		"code": bson.M{
			"$exists": true,
		},
	}, options.Find().SetProjection(projection))
	if err != nil {
		log.WithError(err).Error("Error while finding")
	}

	var fields []interface{}
	for cursor.Next(nil) {
		tmp := bson.M{}
		if err := cursor.Decode(&tmp); err != nil {
			log.WithError(err).Error("Error while cursor iterate")
		}

		fields = append(fields, tmp["locale"])
	}

	return fields
}

func asStrings(target []interface{}) ([]string, error) {
	tmp := []string{}

	for _, val := range target {
		if v, ok := val.(string); ok {
			tmp = append(tmp, v)
			continue
		}
		return nil, errors.New("can't convert elements")
	}

	return tmp, nil
}

func contain(target []string, desired string) bool {
	for _, str := range target {
		if str == desired {
			return true
		}
	}

	return false
}
