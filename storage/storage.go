package storage

import (
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"strings"

	"github.com/mongodb/mongo-go-driver/bson"
	"gitlab.com/project2019-02/thesaurus/models"
)

type Storage interface {
	Raw() interface{}

	StoreDocument(in bson.M) (models.Document, error)
	UpdateDocument(in bson.M, inChanges bson.M) (models.Document, error)
	CountDocuments(t string) (int64, error)
	GetDocuments(opts *options.FindOptions, in bson.M) ([]models.Document, error)
	GetDocument(in bson.M) (models.Document, error)
	DeleteDocuments(in bson.M) (bool, error)
	DeleteDocument(in bson.M) (bool, error)

	StoreOrUpdateMetadata(collectionName string, in *models.UpdMetadata) (models.Metadata, error)
	StoreMetadata(collectionName string, t models.MetadataType) (models.Metadata, error)
	UpdateMetadata(collectionName string, in models.UpdMetadata) (models.Metadata, error)
	GetMetadata(collectionName string) (models.Metadata, error)
	DeleteMetadata(collectionName string) (bool, error)
}

func ValidateCollectionName(name string) string {
	return strings.TrimLeft(name, "_")
}

func ExtractDocumentBsonWithoutType(in bson.M) bson.M {
	m := bson.M{}

	for k, v := range in {
		if k == "type" {
			continue
		}

		m[k] = v
	}

	return m
}

func ExtractDocumentBsonID(in bson.M) bson.M {
	return bson.M{
		"type":   in["type"],
		"locale": in["locale"],
		"code":   in["code"],
	}
}

func ValidateDocumentBsonID(in bson.M) bson.M {
	m := bson.M{
		"locale": in["locale"],
	}

	if _, ok := in["code"]; ok {
		m["code"] = in["code"]
	}

	return m
}
