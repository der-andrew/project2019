package tests

import (
	"os"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/models"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/service"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/tests/integration"
)

var Service *service.Service

type DocumentsSuite struct {
	suite.Suite
}

func (s *DocumentsSuite) SetupTest() {
	_, err := Service.DeleteMetadata("patient.genders_serviceDocumentTest")
	if err != storage.ErrMetadataNotFound {
		s.Require().NoError(err)
	}

	_, err = Service.DeleteDocuments(bson.M{
		"type":   "patient.genders_serviceDocumentTest",
		"locale": "ENG",
	})
	if err != storage.ErrDocumentNotFound {
		s.Require().NoError(err)
	}
}

func (s *DocumentsSuite) TestStoreDocument() {
	item := map[string]interface{}{
		"type":   "patient.genders_serviceDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	_, err := Service.StoreDocument(item)
	s.NoError(err)

	static := models.TypeStatic
	_, err = Service.UpdateMetadata("patient.genders_serviceDocumentTest", models.UpdMetadata{
		Type: &static,
	})
	s.NoError(err)

	item["code"] = "FEMALE"

	_, err = Service.StoreDocument(item)
	s.Error(err)
	s.Equal(err, service.ErrAccessDenied)
}

func (s *DocumentsSuite) TestUpdateDocument() {
	item := map[string]interface{}{
		"type":   "patient.genders_serviceDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	_, err := Service.StoreDocument(item)
	s.NoError(err)

	changes := bson.M{
		"text": "another text",
	}

	doc, err := Service.UpdateDocument(item, changes)
	s.NoError(err)
	s.Equal("another text", *doc.Text)

	static := models.TypeStatic
	_, err = Service.UpdateMetadata("patient.genders_serviceDocumentTest", models.UpdMetadata{
		Type: &static,
	})
	s.NoError(err)

	changes["text"] = "another another text"

	doc, err = Service.UpdateDocument(item, changes)
	s.Error(err)
	s.Equal(service.ErrAccessDenied, err)
}

func (s *DocumentsSuite) TestDeleteDocument() {
	item := map[string]interface{}{
		"type":   "patient.genders_serviceDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	_, err := Service.StoreDocument(item)
	s.NoError(err)

	_, err = Service.DeleteDocument(item)
	s.NoError(err)

	_, err = Service.StoreDocument(item)
	s.NoError(err)

	static := models.TypeStatic
	_, err = Service.UpdateMetadata("patient.genders_serviceDocumentTest", models.UpdMetadata{
		Type: &static,
	})
	s.NoError(err)

	_, err = Service.DeleteDocument(item)
	s.Error(err)
	s.Equal(service.ErrAccessDenied, err)

	dynamic := models.TypeDynamic
	_, err = Service.UpdateMetadata("patient.genders_serviceDocumentTest", models.UpdMetadata{
		Type: &dynamic,
	})
	s.NoError(err)

	_, err = Service.DeleteDocument(item)
	s.NoError(err)

	_, err = Service.GetMetadata("patient.genders_serviceDocumentTest")
	s.Error(err)
	s.Equal(storage.ErrMetadataNotFound, err)
}

func (s *DocumentsSuite) TestDeleteDocuments() {
	item := map[string]interface{}{
		"type":   "patient.genders_serviceDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	_, err := Service.StoreDocument(item)
	s.NoError(err)

	_, err = Service.DeleteDocuments(item)
	s.NoError(err)

	_, err = Service.StoreDocument(item)
	s.NoError(err)

	static := models.TypeStatic
	_, err = Service.UpdateMetadata("patient.genders_serviceDocumentTest", models.UpdMetadata{
		Type: &static,
	})
	s.NoError(err)

	_, err = Service.DeleteDocuments(item)
	s.Error(err)
	s.Equal(service.ErrAccessDenied, err)

	dynamic := models.TypeDynamic
	_, err = Service.UpdateMetadata("patient.genders_serviceDocumentTest", models.UpdMetadata{
		Type: &dynamic,
	})
	s.NoError(err)

	_, err = Service.DeleteDocuments(item)
	s.NoError(err)

	_, err = Service.GetMetadata("patient.genders_serviceDocumentTest")
	s.Error(err)
	s.Equal(storage.ErrMetadataNotFound, err)
}

func TestMain(m *testing.M) {
	// we don't need extra logs in tests
	logrus.SetLevel(logrus.FatalLevel)

	integration.Init() // we need to init env vars and other things, i.e. to know the host and port to connect to db
	integration.Cmd.Run = func(cmd *cobra.Command, args []string) {
		integration.InitService()

		Service = integration.Service

		os.Exit(m.Run())
	}

	if err := integration.Cmd.Execute(); err != nil {
		logrus.Fatal(err.Error())
	}
}

func TestDocuments(t *testing.T) {
	suite.Run(t, new(DocumentsSuite))
}
