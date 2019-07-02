package tests

import (
	"os"
	"testing"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/csv"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/service"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage/mongo"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/tests/integration"
)

var Service *service.Service

type CSVSuite struct {
	suite.Suite
}

func (s *CSVSuite) SetupTest() {
	_, err := Service.DeleteMetadata("patient.genders")
	if err != storage.ErrMetadataNotFound {
		s.Require().NoError(err)
	}

	_, err = Service.DeleteDocuments(bson.M{
		"type":   "patient.genders",
		"locale": "ENG",
	})
	if err != storage.ErrDocumentNotFound {
		s.Require().NoError(err)
	}
}

func (s *CSVSuite) TestImportCSV() {
	path := "../../test_dictionaries"

	csv.ImportCSV(Service.Raw().(*mongo.RawMongo).Database, path, ",")

	for _, item := range []bson.M{
		{
			"type":   "patient.genders",
			"locale": "ENG",
			"code":   "MALE",
			"symbol": "m",
			"text":   "male",
		},
		{
			"type":   "patient.genders",
			"locale": "ENG",
			"code":   "FEMALE",
			"symbol": "f",
			"text":   "female",
		},
		{
			"type":   "patient.genders",
			"locale": "FRA",
			"code":   "MALE",
			"symbol": "m",
			"text":   "mâle",
		},
		{
			"type":   "va.sides",
			"locale": "ENG",
			"code":   "LEFT",
			"text":   "left",
		},
		{
			"type":   "va.sides",
			"locale": "ENG",
			"code":   "RIGHT",
			"text":   "right",
		},
	} {
		filter := storage.ExtractDocumentBsonID(item)

		_, err := Service.GetDocument(filter)
		s.NoError(err, "document with filter %v not found", filter)
	}
}

func TestMain(m *testing.M) {
	// we don't need extra logs in tests
	logrus.SetLevel(logrus.FatalLevel)

	integration.Init() // we need to init env vars and other things, i.e. to know the host and port to connect to db
	integration.Cmd.Run = func(cmd *cobra.Command, args []string) {
		integration.InitService()

		Service = service.Instance

		os.Exit(m.Run())
	}

	if err := integration.Cmd.Execute(); err != nil {
		logrus.Fatal(err.Error())
	}
}

func TestCSV(t *testing.T) {
	suite.Run(t, new(CSVSuite))
}
