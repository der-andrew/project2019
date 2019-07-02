package mongo

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/models"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"
)

func (s *Storage) StoreOrUpdateMetadata(collectionName string, in *models.UpdMetadata) (models.Metadata, error) {
	collection := s.Database.Collection("_metadata")

	filter := bson.M{
		"collectionName": collectionName,
	}

	cur, err := collection.Find(nil, filter)
	if err != nil {
		logrus.WithError(err).Error("Metadata collection finding error")
		return models.Metadata{}, err
	}
	defer func() {
		err := cur.Close(nil)
		if err != nil {
			logrus.WithError(err).Error("Cursor closing error")
		}
	}()

	if cur.Next(nil) {
		if in != nil {
			return s.UpdateMetadata(collectionName, *in)
		}

		return s.UpdateMetadata(collectionName, models.UpdMetadata{
			DateOfChange: models.PtrT(time.Now()),
		})
	}

	return s.StoreMetadata(collectionName, models.DefaultMetadataType)
}

func (s *Storage) StoreMetadata(collectionName string, t models.MetadataType) (models.Metadata, error) {
	collection := s.Database.Collection("_metadata")

	_, err := collection.InsertOne(nil, bson.M{
		"collectionName": collectionName,
		"type":           string(t),
		"dateOfChange":   time.Now().Format(time.RFC3339),
	})
	if err != nil {
		logrus.WithError(err).Error("Metadata collection inserting error")
		return models.Metadata{}, err
	}

	metadata, err := s.GetMetadata(collectionName)
	if err != nil {
		return models.Metadata{}, err
	}

	return metadata, nil
}

func (s *Storage) UpdateMetadata(collectionName string, in models.UpdMetadata) (models.Metadata, error) {
	collection := s.Database.Collection("_metadata")

	filter := bson.M{
		"collectionName": collectionName,
	}

	changes := bson.M{}

	if in.Type != nil {
		changes["type"] = string(*in.Type)
	}

	if in.DateOfChange != nil {
		changes["dateOfChange"] = in.DateOfChange.Format(time.RFC3339)
	}

	_, err := collection.UpdateOne(nil, filter, bson.M{
		"$set": changes,
	})
	if err != nil {
		logrus.WithError(err).Error("Metadata collection finding error")
		return models.Metadata{}, err
	}

	metadata, err := s.GetMetadata(collectionName)
	if err != nil {
		return models.Metadata{}, err
	}

	return metadata, nil
}

func (s *Storage) GetMetadata(collectionName string) (models.Metadata, error) {
	collection := s.Database.Collection("_metadata")

	filter := bson.M{
		"collectionName": collectionName,
	}

	cur, err := collection.Find(nil, filter)
	if err != nil {
		logrus.WithError(err).Error("Metadata collection finding error")
		return models.Metadata{}, err
	}
	defer func() {
		err := cur.Close(nil)
		if err != nil {
			logrus.WithError(err).Error("Cursor closing error")
		}
	}()

	document := bson.M{}
	if cur.Next(nil) {
		if err := cur.Decode(&document); err != nil {
			logrus.WithError(err).Error("Metadata cursor decoding error")
			return models.Metadata{}, err
		}

		if err := cur.Err(); err != nil {
			logrus.WithError(err).Error("Unexpected cursor error")
			return models.Metadata{}, err
		}

		return models.Metadata{
			CollectionName: document["collectionName"].(string),
			Type:           models.MetadataType(document["type"].(string)),
			DateOfChange:   document["dateOfChange"].(string),
		}, nil
	}

	return models.Metadata{}, storage.ErrMetadataNotFound
}

func (s *Storage) DeleteMetadata(collectionName string) (bool, error) {
	collection := s.Database.Collection("_metadata")

	filter := bson.M{
		"collectionName": collectionName,
	}

	_, err := collection.DeleteOne(nil, filter)
	if err != nil {
		logrus.WithError(err).Error("Metadata collection delete error")
		return false, err
	}

	return true, nil
}
