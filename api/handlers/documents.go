package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/models"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/restapi/operations/document"
	dbModels "repo.nefrosovet.ru/maximus-platform/thesaurus/models"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/service"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"
)

func DocumentCreate(params document.DocumentCreateParams) middleware.Responder {
	responseInternalServerError := func(logMessage string) middleware.Responder {
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   logMessage,
		}).Error(InternalServerErrorMessage)

		payload := new(document.DocumentCreateInternalServerErrorBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = &InternalServerErrorMessage
		payload.Errors = nil

		return document.NewDocumentCreateInternalServerError().WithPayload(payload)
	}

	responseMethodNotAllowed := func() middleware.Responder {
		payload := new(document.DocumentCreateMethodNotAllowedBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = MethodNotAllowedMessage("POST")
		payload.Errors = nil

		return document.NewDocumentCreateMethodNotAllowed().WithPayload(payload)
	}

	responseBadRequest := func() middleware.Responder {
		payload := new(document.DocumentCreateBadRequestBody)
		payload.Version = &Version
		payload.Message = PayloadValidationErrorMessage
		payload.Data = nil

		payload.Errors = &models.Document400DataAO1Errors{
			Validation: &models.Document400DataAO1ErrorsValidation{
				Type: "unique",
			},
		}

		return document.NewDocumentCreateBadRequest().WithPayload(payload)
	}

	responseSuccess := func(doc dbModels.Document) middleware.Responder {
		payload := new(document.DocumentCreateOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage
		payload.Errors = nil

		payload.Data = []*document.DataItems0{
			{
				DocumentParam: doc.DocumentParam,
			},
		}

		return document.NewDocumentCreateOK().WithPayload(payload)
	}

	bsonDoc := (&dbModels.Document{
		DocumentParam: *params.Body,
	}).Bson()

	doc, err := Service.StoreDocument(bsonDoc)
	if err != nil {
		switch err {
		case service.ErrAccessDenied:
			return responseMethodNotAllowed()
		case storage.ErrDocumentAlreadyExists:
			return responseBadRequest()
		default:
			return responseInternalServerError(err.Error())
		}
	}

	return responseSuccess(doc)
}

func DocumentCollection(params document.DocumentCollectionParams) middleware.Responder {
	responseInternalServerError := func(logMessage string) middleware.Responder {
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   logMessage,
		}).Error(InternalServerErrorMessage)

		payload := new(document.DocumentCollectionInternalServerErrorBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = &InternalServerErrorMessage
		payload.Errors = nil

		return document.NewDocumentCollectionInternalServerError().WithPayload(payload)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(document.DocumentCollectionNotFoundBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = &NotFoundMessage
		payload.Errors = nil

		return document.NewDocumentCollectionNotFound().WithPayload(payload)
	}

	responseSuccess := func(docs []dbModels.Document) middleware.Responder {
		payload := new(document.DocumentCollectionOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage
		payload.Errors = nil

		for _, doc := range docs {
			payload.Data = append(payload.Data, &document.DataItems0{
				DocumentParam: doc.DocumentParam,
			})
		}

		return document.NewDocumentCollectionOK().WithPayload(payload)
	}

	opts := options.Find()
	if params.Offset != nil {
		opts.SetSkip(int64(*params.Offset)) // Number of documents which should be skip
	}
	if params.Limit != nil {
		opts.SetLimit(int64(*params.Limit)) // Result limit
	}
	opts.SetSort(bsonx.Doc{ // Set crawling order
		{"code", bsonx.Int32(1)},
		{"locale", bsonx.Int32(1)},
	})

	filter := bson.M{
		"type":   params.Type,
		"locale": params.Locale,
	}

	if params.Code != nil {
		filter["code"] = *params.Code
	}

	docs, err := Service.GetDocuments(opts, filter)
	if err != nil {
		return responseInternalServerError(err.Error())
	}

	if params.Code != nil {
		switch {
		case len(docs) == 0:
			return responseNotFound()
		case len(docs) > 1:
			err := EntitiesCountError{
				CollectionName: "Documents",
				Filter:         filter,
				ExpectedCount:  1,
				ActualCount:    len(docs),
			}

			return responseInternalServerError(err.Error())
		}
	}

	return responseSuccess(docs)
}

func DocumentUpdate(params document.DocumentUpdateParams) middleware.Responder {
	responseInternalServerError := func(logMessage string) middleware.Responder {
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   logMessage,
		}).Error(InternalServerErrorMessage)

		payload := new(document.DocumentUpdateInternalServerErrorBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = &InternalServerErrorMessage
		payload.Errors = nil

		return document.NewDocumentUpdateInternalServerError().WithPayload(payload)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(document.DocumentUpdateNotFoundBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = &NotFoundMessage
		payload.Errors = nil

		return document.NewDocumentUpdateNotFound().WithPayload(payload)
	}

	responseMethodNotAllowed := func() middleware.Responder {
		payload := new(document.DocumentUpdateMethodNotAllowedBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = MethodNotAllowedMessage("PUT")
		payload.Errors = nil

		return document.NewDocumentUpdateMethodNotAllowed().WithPayload(payload)
	}

	responseSuccess := func(doc dbModels.Document) middleware.Responder {
		payload := new(document.DocumentUpdateOKBody)
		payload.Version = &Version
		payload.Message = &PayloadSuccessMessage
		payload.Errors = nil

		payload.Data = []*document.DataItems0{
			{
				DocumentParam: doc.DocumentParam,
			},
		}

		return document.NewDocumentUpdateOK().WithPayload(payload)
	}

	filter := (&dbModels.Document{
		DocumentParam: *params.Body,
	}).BsonID()

	changes := (&dbModels.Document{
		DocumentParam: *params.Body,
	}).Bson()

	doc, err := Service.UpdateDocument(filter, changes)
	if err != nil {
		switch err {
		case service.ErrAccessDenied:
			return responseMethodNotAllowed()
		case storage.ErrDocumentNotFound:
			return responseNotFound()
		case storage.ErrDocumentNotModified:
			return responseMethodNotAllowed()
		default:
			return responseInternalServerError(err.Error())
		}
	}

	doc, err = Service.GetDocument(filter)
	if err != nil {
		return responseInternalServerError(err.Error())
	}

	return responseSuccess(doc)
}

func DocumentDelete(params document.DocumentDeleteParams) middleware.Responder {
	responseInternalServerError := func(logMessage string) middleware.Responder {
		logrus.WithFields(logrus.Fields{
			"context": "API",
			"error":   logMessage,
		}).Error(InternalServerErrorMessage)

		payload := new(document.DocumentDeleteInternalServerErrorBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = &InternalServerErrorMessage
		payload.Errors = nil

		return document.NewDocumentDeleteInternalServerError().WithPayload(payload)
	}

	responseNotFound := func() middleware.Responder {
		payload := new(document.DocumentDeleteNotFoundBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = &NotFoundMessage
		payload.Errors = nil

		return document.NewDocumentDeleteNotFound().WithPayload(payload)
	}

	responseNotAllowed := func() middleware.Responder {
		payload := new(document.DocumentDeleteMethodNotAllowedBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = MethodNotAllowedMessage("DELETE")
		payload.Errors = nil

		return document.NewDocumentDeleteMethodNotAllowed().WithPayload(payload)
	}

	responseSuccess := func() middleware.Responder {
		payload := new(document.DocumentDeleteOKBody)
		payload.Version = &Version
		payload.Data = nil
		payload.Message = &PayloadSuccessMessage
		payload.Errors = nil

		return document.NewDocumentDeleteOK().WithPayload(payload)
	}

	filter := bson.M{
		"type": params.Type,
	}

	if params.Locale != nil {
		filter["locale"] = *params.Locale
	}

	if params.Code != nil {
		filter["code"] = *params.Code
	}

	_, err := Service.DeleteDocuments(filter)
	if err != nil {
		switch err {
		case service.ErrAccessDenied:
			return responseNotAllowed()
		case storage.ErrDocumentNotFound:
			return responseNotFound()
		default:
			return responseInternalServerError(err.Error())
		}
	}

	return responseSuccess()
}
