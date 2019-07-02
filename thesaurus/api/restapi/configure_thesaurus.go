// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/handlers"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/restapi/operations/document"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/restapi/operations/status"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/service"
)

//go:generate swagger generate server --target ../../api --name Thesaurus --spec ../../docs/swagger.yaml --exclude-main

// Version of service
var Version string

func configureFlags(api *operations.ThesaurusAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ThesaurusAPI) http.Handler {
	handlers.Version = Version
	handlers.Service = service.Instance

	api.Logger = logrus.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.StatusStatusViewHandler = status.StatusViewHandlerFunc(handlers.StatusView)

	api.DocumentDocumentCreateHandler = document.DocumentCreateHandlerFunc(handlers.DocumentCreate)

	api.DocumentDocumentCollectionHandler = document.DocumentCollectionHandlerFunc(handlers.DocumentCollection)

	api.DocumentDocumentUpdateHandler = document.DocumentUpdateHandlerFunc(handlers.DocumentUpdate)

	api.DocumentDocumentDeleteHandler = document.DocumentDeleteHandlerFunc(handlers.DocumentDelete)

	api.ServerShutdown = func() {}

	api.ServeError = handlers.ServeError

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}