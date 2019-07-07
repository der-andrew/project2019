package tests

import (
	"os"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"

	"gitlab-host/maximus-platform/thesaurus/storage"
	"gitlab-host/maximus-platform/thesaurus/tests/integration"
)

var Storage storage.Storage

type DocumentsSuite struct {
	suite.Suite
}

func (s *DocumentsSuite) SetupTest() {
	_, err := Storage.DeleteDocuments(bson.M{
		"type":   "patient.genders_storageDocumentTest",
		"locale": "ENG",
	})
	if err != storage.ErrDocumentNotFound {
		s.Require().NoError(err)
	}
}

func (s *DocumentsSuite) TestStoreDocument() {
	item := map[string]interface{}{
		"type":   "patient.genders_storageDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	doc, err := Storage.StoreDocument(item)
	s.NoError(err)

	_, err = Storage.StoreDocument(item)
	s.Error(err)
	s.Equal(storage.ErrDocumentAlreadyExists, err)

	foundDoc, err := Storage.GetDocument(item)
	s.NoError(err)

	s.Equal(foundDoc, doc)
}

func (s *DocumentsSuite) TestUpdateDocument() {
	item := map[string]interface{}{
		"type":   "patient.genders_storageDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	changes := bson.M{
		"text": "another text",
	}

	_, err := Storage.UpdateDocument(item, changes)
	s.Error(err)
	s.Equal(storage.ErrDocumentNotFound, err)

	_, err = Storage.StoreDocument(item)
	s.NoError(err)

	doc, err := Storage.UpdateDocument(item, changes)
	s.NoError(err)

	foundDoc, err := Storage.GetDocument(item)
	s.NoError(err)

	s.Equal(foundDoc, doc)

	_, err = Storage.UpdateDocument(item, changes)
	s.Error(err)
	s.Equal(storage.ErrDocumentNotModified, err)
}

func (s *DocumentsSuite) TestCountDocuments() {
	item := map[string]interface{}{
		"type":   "patient.genders_storageDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	count, err := Storage.CountDocuments("patient.genders_storageDocumentTest")
	s.NoError(err)
	s.Equal(int64(0), count)

	_, err = Storage.StoreDocument(item)
	s.NoError(err)

	count, err = Storage.CountDocuments("patient.genders_storageDocumentTest")
	s.NoError(err)
	s.Equal(int64(1), count)
}

func (s *DocumentsSuite) TestGetDocuments() {
	item := map[string]interface{}{
		"type":   "patient.genders_storageDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	docs, err := Storage.GetDocuments(nil, item)
	s.NoError(err)
	s.Len(docs, 0)

	_, err = Storage.StoreDocument(item)
	s.NoError(err)

	item["code"] = "FEMALE"

	_, err = Storage.StoreDocument(item)
	s.NoError(err)

	delete(item, "code")

	docs, err = Storage.GetDocuments(nil, item)
	s.NoError(err)
	s.Len(docs, 2)
}

func (s *DocumentsSuite) TestGetDocument() {
	item := map[string]interface{}{
		"type":   "patient.genders_storageDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	_, err := Storage.GetDocument(item)
	s.Error(err)
	s.Equal(storage.ErrDocumentNotFound, err)

	_, err = Storage.StoreDocument(item)
	s.NoError(err)

	doc, err := Storage.GetDocument(item)
	s.NoError(err)

	s.Equal(item["type"].(string), *doc.Type)
	s.Equal(item["locale"].(string), *doc.Locale)
	s.Equal(item["code"].(string), *doc.Code)
	s.Equal(item["text"].(string), *doc.Text)
}

func (s *DocumentsSuite) TestDeleteDocument() {
	item := map[string]interface{}{
		"type":   "patient.genders_storageDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	_, err := Storage.DeleteDocument(item)
	s.Error(err)
	s.Equal(storage.ErrDocumentNotFound, err)

	_, err = Storage.StoreDocument(item)
	s.NoError(err)

	_, err = Storage.DeleteDocument(item)
	s.NoError(err)

	docs, err := Storage.GetDocuments(nil, item)
	s.NoError(err)
	s.Len(docs, 0)
}

func (s *DocumentsSuite) TestDeleteDocuments() {
	item := map[string]interface{}{
		"type":   "patient.genders_storageDocumentTest",
		"locale": "ENG",
		"code":   "MALE",
		"text":   "some text",
	}

	_, err := Storage.DeleteDocument(item)
	s.Error(err)
	s.Equal(storage.ErrDocumentNotFound, err)

	_, err = Storage.StoreDocument(item)
	s.NoError(err)

	item["code"] = "FEMALE"

	_, err = Storage.StoreDocument(item)
	s.NoError(err)

	delete(item, "code")

	_, err = Storage.DeleteDocuments(item)
	s.NoError(err)

	docs, err := Storage.GetDocuments(nil, item)
	s.NoError(err)
	s.Len(docs, 0)
}

func TestMain(m *testing.M) {
	// we don't need extra logs in tests
	logrus.SetLevel(logrus.FatalLevel)

	integration.Init() // we need to init env vars and other things, i.e. to know the host and port to connect to db
	integration.Cmd.Run = func(cmd *cobra.Command, args []string) {
		integration.InitStorage()

		Storage = integration.Storage

		os.Exit(m.Run())
	}

	if err := integration.Cmd.Execute(); err != nil {
		logrus.Fatal(err.Error())
	}
}

func TestDocuments(t *testing.T) {
	suite.Run(t, new(DocumentsSuite))
}
