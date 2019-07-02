package mongo

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/models"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"
)

func (s *Storage) StoreDocument(in bson.M) (models.Document, error) {
	_, err := s.GetDocument(storage.ExtractDocumentBsonID(in))
	switch err {
	case storage.ErrDocumentNotFound:
	case nil:
		return models.Document{}, storage.ErrDocumentAlreadyExists
	default:
		return models.Document{}, err
	}

	collection := s.Database.Collection(
		storage.ValidateCollectionName(in["type"].(string)),
	)

	_, err = collection.InsertOne(nil, storage.ExtractDocumentBsonWithoutType(in))
	if err != nil {
		logrus.WithError(err).Error("Document collection inserting error")
		return models.Document{}, err
	}

	return models.DocumentFromBson(in), err
}

func (s *Storage) UpdateDocument(in bson.M, inChanges bson.M) (models.Document, error) {
	_, err := s.GetDocument(storage.ExtractDocumentBsonID(in))
	if err != nil {
		return models.Document{}, err
	}

	collection := s.Database.Collection(
		storage.ValidateCollectionName(in["type"].(string)),
	)

	filter := storage.ValidateDocumentBsonID(in)

	changes := bson.M{
		"$set": storage.ExtractDocumentBsonWithoutType(inChanges),
	}

	res, err := collection.UpdateOne(nil, filter, changes)
	if err != nil {
		logrus.WithError(err).Error("Document collection updating error")
		return models.Document{}, err
	}

	if res.ModifiedCount == 0 {
		return models.Document{}, storage.ErrDocumentNotModified
	}

	doc, err := s.GetDocument(in)
	if err != nil {
		return models.Document{}, err
	}

	return doc, err
}

func (s *Storage) CountDocuments(t string) (int64, error) {
	collection := s.Database.Collection(
		storage.ValidateCollectionName(t),
	)

	count, err := collection.Count(nil, bson.M{})
	if err != nil {
		logrus.WithError(err).Error("Document count error")
		return -1, err
	}

	return count, nil
}

func (s *Storage) GetDocuments(opts *options.FindOptions, in bson.M) ([]models.Document, error) {
	collection := s.Database.Collection(
		storage.ValidateCollectionName(in["type"].(string)),
	)

	filter := storage.ValidateDocumentBsonID(in)

	cur, err := collection.Find(nil, filter, opts)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = cur.Close(nil); err != nil {
			logrus.WithError(err).Error("Cursor closing error")
		}
	}()

	// Get next and decoded docs
	var doc *bson.M
	var docs []models.Document
	for cur.Next(context.TODO()) {
		err := cur.Decode(&doc)
		if err != nil {
			logrus.WithError(err).Error("Cursor decoding error")
			return nil, err
		}

		if doc != nil {
			delete(*doc, "_id")
			(*doc)["type"] = collection.Name()

			docs = append(docs, models.DocumentFromBson(*doc))

			doc = nil
		}
	}

	if err := cur.Err(); err != nil {
		logrus.WithError(err).Error("Unexpected cursor error")
		return nil, err
	}

	return docs, nil
}

func (s *Storage) GetDocument(in bson.M) (models.Document, error) {
	collection := s.Database.Collection(
		storage.ValidateCollectionName(in["type"].(string)),
	)

	filter := storage.ValidateDocumentBsonID(in)

	cur, err := collection.Find(nil, filter)
	if err != nil {
		logrus.WithError(err).Error("Document collection finding error")
		return models.Document{}, err
	}
	defer func() {
		if err = cur.Close(nil); err != nil {
			logrus.WithError(err).Error("Cursor closing error")
		}
	}()

	document := bson.M{
		"type": collection.Name(),
	}

	if cur.Next(nil) {
		if err := cur.Decode(&document); err != nil {
			logrus.WithError(err).Error("Cursor decoding error")
			return models.Document{}, err
		}

		delete(document, "_id")

		if err := cur.Err(); err != nil {
			logrus.WithError(err).Error("Unexpected cursor error")
			return models.DocumentFromBson(document), err
		}

		if _, ok := document["code"]; ok {
			return models.DocumentFromBson(document), nil
		}
	}

	return models.Document{}, storage.ErrDocumentNotFound
}

func (s *Storage) DeleteDocument(in bson.M) (bool, error) {
	collection := s.Database.Collection(
		storage.ValidateCollectionName(in["type"].(string)),
	)

	filter := storage.ValidateDocumentBsonID(in)

	res, err := collection.DeleteOne(nil, filter)
	if err != nil {
		return false, err
	}

	if res.DeletedCount == 0 {
		return false, storage.ErrDocumentNotFound
	}

	count, err := collection.Count(nil, bson.M{})
	if err != nil {
		return true, err
	}

	if count == 0 {
		if err := collection.Drop(nil); err != nil {
			return true, err
		}
	}

	return true, nil
}

func (s *Storage) DeleteDocuments(in bson.M) (bool, error) {
	collection := s.Database.Collection(
		storage.ValidateCollectionName(in["type"].(string)),
	)

	filter := storage.ValidateDocumentBsonID(in)

	if filter["locale"] == nil && filter["code"] == nil {
		err := collection.Drop(nil)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	res, err := collection.DeleteMany(nil, filter)
	if err != nil {
		return false, err
	}

	if res.DeletedCount == 0 {
		return false, storage.ErrDocumentNotFound
	}

	count, err := collection.Count(nil, bson.M{})
	if err != nil {
		return true, err
	}

	if count == 0 {
		if err := collection.Drop(nil); err != nil {
			return true, err
		}
	}

	return true, nil
}
