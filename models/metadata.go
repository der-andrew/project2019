package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

type MetadataType string

const (
	TypeStatic  MetadataType = "STATIC"
	TypeDynamic MetadataType = "DYNAMIC"
)

var DefaultMetadataType MetadataType = TypeDynamic

type Metadata struct {
	CollectionName string
	Type           MetadataType

	DateOfChange string
}

type UpdMetadata struct {
	Type         *MetadataType
	DateOfChange *time.Time
}

func MetadataFromBson(m bson.M) Metadata {
	metadata := Metadata{}

	for k, v := range m {
		switch k {
		case "collectionName":
			metadata.CollectionName = v.(string)
		case "type":
			metadata.Type = MetadataType(v.(string))
		case "dateOfChange":
			metadata.DateOfChange = v.(string)
		}
	}

	return metadata
}
