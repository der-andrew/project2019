package models

import (
	"github.com/mongodb/mongo-go-driver/bson"

	"gitlab.com/project2019-02/thesaurus/api/models"
)

type Document struct {
	models.DocumentParam
}

func (d *Document) BsonID() bson.M {
	m := bson.M{
		"type":   *d.Type,
		"locale": *d.Locale,
		"code":   *d.Code,
	}

	return m
}

func (d *Document) Bson() bson.M {
	m := bson.M{
		"type":   *d.Type,
		"locale": *d.Locale,
		"code":   *d.Code,
		"text":   *d.Text,
	}

	for k, v := range d.DocumentParamAdditionalProperties {
		m[k] = v
	}

	return m
}

func DocumentFromBson(m bson.M) Document {
	dm := models.DocumentParam{
		DocumentParamAdditionalProperties: map[string]interface{}{},
	}

	for k, v := range m {
		switch k {
		case "type":
			dm.Type = PtrS(m["type"].(string))
		case "locale":
			dm.Locale = PtrS(m["locale"].(string))
		case "code":
			dm.Code = PtrS(m["code"].(string))
		case "text":
			dm.Text = PtrS(m["text"].(string))
		default:
			dm.DocumentParamAdditionalProperties[k] = v
		}
	}

	return Document{
		DocumentParam: dm,
	}
}
