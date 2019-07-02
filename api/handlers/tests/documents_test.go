package tests

import (
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/handlers"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/models"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/restapi/operations/document"
	dbModels "repo.nefrosovet.ru/maximus-platform/thesaurus/models"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"
)

type DocumentsSuite struct {
	suite.Suite
}

func (s *DocumentsSuite) SetupTest() {
	_, err := Service.DeleteMetadata("patient.genders_handlersDocumentTest")
	if err != storage.ErrMetadataNotFound {
		s.Require().NoError(err)
	}

	for _, locale := range []string{"ENG", "FRA", "RUS"} {
		_, err = Service.DeleteDocuments(bson.M{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": locale,
		})
		if err != storage.ErrDocumentNotFound {
			s.Require().NoError(err)
		}
	}
}

func (s *DocumentsSuite) TestDocumentCreate() {
	// Prepare

	defaultParams := func() document.DocumentCreateParams {
		params := document.NewDocumentCreateParams()
		params.Body = new(models.DocumentParam)

		params.Body.Type = dbModels.PtrS("patient.genders_handlersDocumentTest")
		params.Body.Locale = dbModels.PtrS("ENG")
		params.Body.Code = dbModels.PtrS("MALE")
		params.Body.Text = dbModels.PtrS("some text")

		params.Body.DocumentParamAdditionalProperties = map[string]interface{}{}

		return params
	}

	verify := func(params document.DocumentCreateParams, typeMustBe interface{}) interface{} {
		res := handlers.DocumentCreate(params)
		s.IsType(typeMustBe, res)

		switch res := res.(type) {
		case *document.DocumentCreateOK:
			checkResponseServiceFields(s.Assert(), res.Payload.SuccessData)
			return res.Payload
		case *document.DocumentCreateBadRequest:
			checkResponseServiceFields(s.Assert(), res.Payload.Document400Data)
			return res.Payload
		case *document.DocumentCreateMethodNotAllowed:
			checkResponseServiceFields(s.Assert(), res.Payload.Error405Data)
			return res.Payload
		default:
			s.T().Errorf("Unknown response type: %T", res)
		}

		return nil
	}

	// Test

	{
		// Post new doc

		params := defaultParams()
		payload := verify(params, &document.DocumentCreateOK{}).(*document.DocumentCreateOKBody)

		s.Len(payload.Data, 1)
		s.Equal(*params.Body, payload.Data[0].DocumentParam)
	}

	{
		// Post exist doc

		params := defaultParams()
		item := map[string]interface{}{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "ENG",
			"code":   "MALE",
			"text":   "some text",
		}

		_, err := Service.StoreDocument(item)
		s.Error(err)
		s.Equal(storage.ErrDocumentAlreadyExists, err)

		verify(params, &document.DocumentCreateBadRequest{})
	}

	{
		// Post doc to private thesaurus
		s.SetupTest()

		params := defaultParams()
		item := map[string]interface{}{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "ENG",
			"code":   "MALE",
			"text":   "some text",
		}

		_, err := Service.StoreDocument(item)
		s.NoError(err)

		static := dbModels.TypeStatic
		_, err = Service.UpdateMetadata(item["type"].(string), dbModels.UpdMetadata{
			Type: &static,
		})
		s.NoError(err)

		params.Body.Locale = dbModels.PtrS("RUS")

		verify(params, &document.DocumentCreateMethodNotAllowed{})
	}
}

func (s *DocumentsSuite) TestDocumentCollection() {
	// Prepare

	defaultParams := func() document.DocumentCollectionParams {
		params := document.NewDocumentCollectionParams()

		params.Type = "patient.genders_handlersDocumentTest"
		params.Locale = "ENG"
		params.Limit = dbModels.PtrI64(10)
		params.Offset = dbModels.PtrI64(0)

		return params
	}

	verify := func(params document.DocumentCollectionParams, typeMustBe interface{}) interface{} {
		res := handlers.DocumentCollection(params)
		s.IsType(typeMustBe, res)

		switch res := res.(type) {
		case *document.DocumentCollectionOK:
			checkResponseServiceFields(s.Assert(), res.Payload.SuccessData)
			return res.Payload
		case *document.DocumentCollectionNotFound:
			checkResponseServiceFields(s.Assert(), res.Payload.Error404Data)
			return res.Payload
		case *document.DocumentCollectionMethodNotAllowed:
			checkResponseServiceFields(s.Assert(), res.Payload.Error405Data)
			return res.Payload
		default:
			s.T().Errorf("Unknown response type: %T", res)
		}

		return nil
	}

	items := []map[string]interface{}{
		{
			"type":      "patient.genders_handlersDocumentTest",
			"locale":    "ENG",
			"code":      "M1",
			"symbol":    "m",
			"text":      "male",
			"protected": true,
		},
		{
			"type":      "patient.genders_handlersDocumentTest",
			"locale":    "ENG",
			"code":      "M2",
			"symbol":    "f",
			"text":      "female",
			"protected": true,
		},
		{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "FRA",
			"code":   "M3",
			"symbol": "m",
			"text":   "mâle",
		},
		{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "ENG",
			"code":   "M4",
			"text":   "left",
		},
		{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "ENG",
			"code":   "M5",
			"text":   "right",
		},
	}

	for _, obj := range items {
		_, err := Service.StoreDocument(obj)
		s.NoError(err)
	}

	// Test

	{
		// Get all

		payload := verify(defaultParams(), &document.DocumentCollectionOK{}).(*document.DocumentCollectionOKBody)

		s.Len(payload.Data, 4)
		s.EqualValues([]*document.DataItems0{
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("ENG"),
					Code:   dbModels.PtrS("M1"),
					Text:   dbModels.PtrS("male"),

					DocumentParamAdditionalProperties: map[string]interface{}{
						"symbol":    "m",
						"protected": true,
					},
				},
			},
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("ENG"),
					Code:   dbModels.PtrS("M2"),
					Text:   dbModels.PtrS("female"),

					DocumentParamAdditionalProperties: map[string]interface{}{
						"symbol":    "f",
						"protected": true,
					},
				},
			},
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("ENG"),
					Code:   dbModels.PtrS("M4"),
					Text:   dbModels.PtrS("left"),

					DocumentParamAdditionalProperties: map[string]interface{}{},
				},
			},
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("ENG"),
					Code:   dbModels.PtrS("M5"),
					Text:   dbModels.PtrS("right"),

					DocumentParamAdditionalProperties: map[string]interface{}{},
				},
			},
		}, payload.Data)
	}

	{
		// Get with code

		params := defaultParams()
		params.Code = dbModels.PtrS("M1")

		payload := verify(params, &document.DocumentCollectionOK{}).(*document.DocumentCollectionOKBody)

		s.Len(payload.Data, 1)
		s.EqualValues([]*document.DataItems0{
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("ENG"),
					Code:   dbModels.PtrS("M1"),
					Text:   dbModels.PtrS("male"),

					DocumentParamAdditionalProperties: map[string]interface{}{
						"symbol":    "m",
						"protected": true,
					},
				},
			},
		}, payload.Data)
	}

	{
		// Get with offset

		params := defaultParams()
		params.Limit = dbModels.PtrI64(2)
		params.Offset = dbModels.PtrI64(2)

		payload := verify(params, &document.DocumentCollectionOK{}).(*document.DocumentCollectionOKBody)

		s.Len(payload.Data, 2)
		s.EqualValues([]*document.DataItems0{
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("ENG"),
					Code:   dbModels.PtrS("M4"),
					Text:   dbModels.PtrS("left"),

					DocumentParamAdditionalProperties: map[string]interface{}{},
				},
			},
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("ENG"),
					Code:   dbModels.PtrS("M5"),
					Text:   dbModels.PtrS("right"),

					DocumentParamAdditionalProperties: map[string]interface{}{},
				},
			},
		}, payload.Data)
	}

	{
		// Get with limit

		params := defaultParams()
		params.Limit = dbModels.PtrI64(2)
		params.Offset = dbModels.PtrI64(0)

		payload := verify(params, &document.DocumentCollectionOK{}).(*document.DocumentCollectionOKBody)

		s.Len(payload.Data, 2)
		s.EqualValues([]*document.DataItems0{
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("ENG"),
					Code:   dbModels.PtrS("M1"),
					Text:   dbModels.PtrS("male"),

					DocumentParamAdditionalProperties: map[string]interface{}{
						"symbol":    "m",
						"protected": true,
					},
				},
			},
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("ENG"),
					Code:   dbModels.PtrS("M2"),
					Text:   dbModels.PtrS("female"),

					DocumentParamAdditionalProperties: map[string]interface{}{
						"symbol":    "f",
						"protected": true,
					},
				},
			},
		}, payload.Data)
	}

	{
		// Get more than exist

		params := defaultParams()
		params.Locale = "FRA"
		params.Limit = dbModels.PtrI64(9000)
		params.Offset = dbModels.PtrI64(0)

		payload := verify(params, &document.DocumentCollectionOK{}).(*document.DocumentCollectionOKBody)

		s.Len(payload.Data, 1)
		s.EqualValues([]*document.DataItems0{
			{
				DocumentParam: models.DocumentParam{
					Type:   dbModels.PtrS("patient.genders_handlersDocumentTest"),
					Locale: dbModels.PtrS("FRA"),
					Code:   dbModels.PtrS("M3"),
					Text:   dbModels.PtrS("mâle"),

					DocumentParamAdditionalProperties: map[string]interface{}{
						"symbol": "m",
					},
				},
			},
		}, payload.Data)
	}
}

func (s *DocumentsSuite) TestDocumentUpdate() {
	// Prepare

	defaultParams := func() document.DocumentUpdateParams {
		params := document.NewDocumentUpdateParams()
		params.Body = new(models.DocumentParam)

		params.Body.Type = dbModels.PtrS("patient.genders_handlersDocumentTest")
		params.Body.Locale = dbModels.PtrS("ENG")
		params.Body.Code = dbModels.PtrS("MALE")
		params.Body.Text = dbModels.PtrS("some text")

		params.Body.DocumentParamAdditionalProperties = map[string]interface{}{}

		return params
	}

	verify := func(params document.DocumentUpdateParams, typeMustBe interface{}) interface{} {
		res := handlers.DocumentUpdate(params)
		s.IsType(typeMustBe, res)

		switch res := res.(type) {
		case *document.DocumentUpdateOK:
			checkResponseServiceFields(s.Assert(), res.Payload.SuccessData)
			return res.Payload
		case *document.DocumentUpdateNotFound:
			checkResponseServiceFields(s.Assert(), res.Payload.Error404Data)
			return res.Payload
		case *document.DocumentUpdateMethodNotAllowed:
			checkResponseServiceFields(s.Assert(), res.Payload.Error405Data)
			return res.Payload
		default:
			s.T().Errorf("Unknown response type: %T", res)
		}

		return nil
	}

	// Test

	{
		// Put exist docs

		params := defaultParams()
		item := map[string]interface{}{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "ENG",
			"code":   "MALE",
			"text":   "some text",
		}

		_, err := Service.StoreDocument(item)
		s.NoError(err)

		params.Body.Text = dbModels.PtrS("new text")

		payload := verify(params, &document.DocumentUpdateOK{}).(*document.DocumentUpdateOKBody)
		s.Len(payload.Data, 1)
		s.Equal(*params.Body, payload.Data[0].DocumentParam)
	}

	{
		// Put non-existent docs
		s.SetupTest()

		verify(defaultParams(), &document.DocumentUpdateNotFound{})
	}

	{
		// Put doc at private thesaurus
		s.SetupTest()

		params := defaultParams()
		item := map[string]interface{}{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "ENG",
			"code":   "MALE",
			"text":   "some text",
		}

		_, err := Service.StoreDocument(item)
		s.NoError(err)

		static := dbModels.TypeStatic
		_, err = Service.UpdateMetadata(item["type"].(string), dbModels.UpdMetadata{
			Type: &static,
		})
		s.NoError(err)

		verify(params, &document.DocumentUpdateMethodNotAllowed{})
	}

	{
		// Put without changes
		s.SetupTest()

		params := defaultParams()
		item := map[string]interface{}{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "ENG",
			"code":   "MALE",
			"text":   "some text",
		}

		_, err := Service.StoreDocument(item)
		s.NoError(err)

		verify(params, &document.DocumentUpdateMethodNotAllowed{})
	}
}

func (s *DocumentsSuite) TestDocumentDelete() {
	// Prepare

	defaultParams := func() document.DocumentDeleteParams {
		params := document.NewDocumentDeleteParams()

		params.Type = "patient.genders_handlersDocumentTest"
		params.Locale = dbModels.PtrS("ENG")
		params.Code = dbModels.PtrS("MALE")

		return params
	}

	verify := func(params document.DocumentDeleteParams, typeMustBe interface{}) interface{} {
		res := handlers.DocumentDelete(params)
		s.IsType(typeMustBe, res)

		switch res := res.(type) {
		case *document.DocumentDeleteOK:
			checkResponseServiceFields(s.Assert(), res.Payload.SuccessData)
			return res.Payload
		case *document.DocumentDeleteNotFound:
			checkResponseServiceFields(s.Assert(), res.Payload.Error404Data)
			return res.Payload
		case *document.DocumentDeleteMethodNotAllowed:
			checkResponseServiceFields(s.Assert(), res.Payload.Error405Data)
			return res.Payload
		default:
			s.T().Errorf("Unknown response type: %T", res)
		}

		return nil
	}

	// Test

	{
		// Delete existing document

		params := defaultParams()
		item := map[string]interface{}{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "ENG",
			"code":   "MALE",
			"text":   "some text",
		}

		_, err := Service.StoreDocument(item)
		s.NoError(err)

		verify(params, &document.DocumentDeleteOK{})

		_, err = Service.GetDocument(item)
		s.Error(err)
		s.Equal(storage.ErrDocumentNotFound, err)
	}

	{
		// Delete non-existent document
		s.SetupTest()

		params := defaultParams()
		params.Code = dbModels.PtrS("FEMALE")

		verify(params, &document.DocumentDeleteNotFound{})
	}

	{
		// Delete protected document (only one doc)
		s.SetupTest()

		params := defaultParams()
		item := map[string]interface{}{
			"type":   "patient.genders_handlersDocumentTest",
			"locale": "ENG",
			"code":   "MALE",
			"text":   "some text",
		}

		_, err := Service.StoreDocument(item)
		s.NoError(err)

		static := dbModels.TypeStatic
		_, err = Service.UpdateMetadata(item["type"].(string), dbModels.UpdMetadata{
			Type: &static,
		})
		s.NoError(err)

		verify(params, &document.DocumentDeleteMethodNotAllowed{})
	}
}

func TestDocuments(t *testing.T) {
	suite.Run(t, new(DocumentsSuite))
}
