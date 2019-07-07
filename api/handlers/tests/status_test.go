package tests

import (
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/stretchr/testify/suite"

	"gitlab-host/maximus-platform/thesaurus/api/handlers"
	"gitlab-host/maximus-platform/thesaurus/api/models"
	"gitlab-host/maximus-platform/thesaurus/api/restapi/operations/status"
	dbModels "gitlab-host/maximus-platform/thesaurus/models"
	"gitlab-host/maximus-platform/thesaurus/storage"
)

type StatusSuite struct {
	suite.Suite
}

func (s *StatusSuite) SetupTest() {
	_, err := Service.DeleteMetadata("patient.genders_handlersStatusTest")
	if err != storage.ErrMetadataNotFound {
		s.Require().NoError(err)
	}

	_, err = Service.DeleteMetadata("va.sides_handlersStatusTest")
	if err != storage.ErrMetadataNotFound {
		s.Require().NoError(err)
	}

	for _, locale := range []string{"ENG", "FRA", "RUS"} {
		_, err = Service.DeleteDocuments(bson.M{
			"type":   "patient.genders_handlersStatusTest",
			"locale": locale,
		})
		if err != storage.ErrDocumentNotFound {
			s.Require().NoError(err)
		}
	}

	for _, locale := range []string{"ENG", "RUS"} {
		_, err = Service.DeleteDocuments(bson.M{
			"type":   "va.sides_handlersStatusTest",
			"locale": locale,
		})
		if err != storage.ErrDocumentNotFound {
			s.Require().NoError(err)
		}
	}

}

func (s *StatusSuite) getUpdateDate(collectionName string) string {
	metadata, err := Service.GetMetadata(collectionName)
	s.NoError(err)

	return metadata.DateOfChange
}

func (s *StatusSuite) TestStatusView() {
	// Prepare

	defaultParams := func() status.StatusViewParams {
		return status.NewStatusViewParams()
	}

	verify := func(params status.StatusViewParams, typeMustBe interface{}) interface{} {
		res := handlers.StatusView(params)
		s.IsType(typeMustBe, res)

		switch res := res.(type) {
		case *status.StatusViewOK:
			checkResponseServiceFields(s.Assert(), res.Payload.SuccessData)
			return res.Payload
		case *status.StatusViewMethodNotAllowed:
			checkResponseServiceFields(s.Assert(), res.Payload.Error405Data)
			return res.Payload
		default:
			s.T().Errorf("Unknown response type: %T", res)
		}

		return nil
	}

	items := []map[string]interface{}{
		{
			"type":   "patient.genders_handlersStatusTest",
			"locale": "ENG",
			"code":   "MALE",
			"text":   "male",
		},
		{
			"type":   "patient.genders_handlersStatusTest",
			"locale": "FRA",
			"code":   "MALE",
			"text":   "mâle",
		},
		{
			"type":   "patient.genders_handlersStatusTest",
			"locale": "RUS",
			"code":   "MALE",
			"text":   "мужской",
		},
		{
			"type":   "va.sides_handlersStatusTest",
			"locale": "ENG",
			"code":   "LEFT",
			"text":   "left",
		},
		{
			"type":   "va.sides_handlersStatusTest",
			"locale": "RUS",
			"code":   "LEFT",
			"text":   "слева",
		},
	}

	for _, obj := range items {
		_, err := Service.StoreDocument(obj)
		s.NoError(err)
	}

	// Test

	{
		// Test provided test set dictionaries

		payload := verify(defaultParams(), &status.StatusViewOK{}).(*status.StatusViewOKBody)

		items := []*status.DataItems0{
			{
				StatusObject: models.StatusObject{
					Name:    "va.sides_handlersStatusTest",
					Locales: []string{"ENG", "RUS"},
					Type:    string(dbModels.TypeDynamic),
					Updated: s.getUpdateDate("va.sides_handlersStatusTest"),
				},
			},
			{
				StatusObject: models.StatusObject{
					Name:    "patient.genders_handlersStatusTest",
					Locales: []string{"ENG", "FRA", "RUS"},
					Type:    string(dbModels.TypeDynamic),
					Updated: s.getUpdateDate("patient.genders_handlersStatusTest"),
				},
			},
		}

		for _, item := range items {
			found := false

			for _, payloadItem := range payload.Data {
				if item.Name == payloadItem.Name {
					s.Equal(item.Locales, payloadItem.Locales)
					s.Equal(item.Type, payloadItem.Type)
					s.Equal(item.Updated, payloadItem.Updated)

					found = true
				}
			}

			s.True(found)
		}
	}
}

func TestStatus(t *testing.T) {
	suite.Run(t, new(StatusSuite))
}
