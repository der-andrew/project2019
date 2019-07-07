package tests

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"gitlab-host/maximus-platform/thesaurus/api/handlers"
	"gitlab-host/maximus-platform/thesaurus/api/models"
	"gitlab-host/maximus-platform/thesaurus/service"
	"gitlab-host/maximus-platform/thesaurus/tests/integration"
)

var Service *service.Service

func TestMain(m *testing.M) {
	// we don't need extra logs in tests
	logrus.SetLevel(logrus.FatalLevel)

	integration.Init() // we need to init env vars and other things, i.e. to know the host and port to connect to db
	integration.Cmd.Run = func(cmd *cobra.Command, args []string) {
		integration.InitService()

		Service = integration.Service
		handlers.Service = integration.Service

		os.Exit(m.Run())
	}

	if err := integration.Cmd.Execute(); err != nil {
		logrus.Fatal(err.Error())
	}
}

func checkResponseServiceFields(r *assert.Assertions, payload interface{}) {
	switch p := payload.(type) {
	case models.SuccessData:
		r.Equal(handlers.PayloadSuccessMessage, *p.Message)
		r.Empty(p.Errors)
		r.Equal(handlers.Version, *p.Version)
	case models.Document400Data:
		r.NotEmpty(p.Errors)
		r.Equal(handlers.Version, *p.Version)
	case models.Error400Data:
		r.Equal(handlers.PayloadValidationErrorMessage, *p.Message)
		r.Empty(p.Errors)
		r.Equal(handlers.Version, *p.Version)
	case models.Error404Data:
		r.Equal(handlers.NotFoundMessage, *p.Message)
		r.Equal(handlers.Version, *p.Version)
	case models.Error405Data:
		r.Equal(handlers.Version, *p.Version)
	case models.Error500Data:
		r.Equal(handlers.InternalServerErrorMessage, *p.Message)
		r.Equal(handlers.Version, *p.Version)
	default:
		logrus.Fatalf("There is no check of service fields for %T", p)
	}
}
