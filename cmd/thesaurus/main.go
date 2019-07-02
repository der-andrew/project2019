package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/getsentry/raven-go"
	"github.com/go-openapi/loads"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/restapi"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/api/restapi/operations"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/csv"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/service"
	"repo.nefrosovet.ru/maximus-platform/thesaurus/storage/mongo"
)

var (
	version = "No Version Provided"
	cfgFile string
)

var cmd = &cobra.Command{
	Use:     "thesaurus",
	Short:   "Thesaurus is a service for documents management",
	Long:    `For getting more information read project wiki`,
	Version: version,
	Run: func(cmd *cobra.Command, args []string) {
		checkMandatoryParams()
		start()
	},
}

func init() {
	restapi.Version = version

	cmd.SetVersionTemplate(fmt.Sprintf("Version: %s\n", restapi.Version))

	cobra.OnInitialize(initConfig)

	/* Set command flags */
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")

	cmd.PersistentFlags().String("http.host", "0.0.0.0", "Host address")
	cmd.PersistentFlags().Int("http.port", 80, "Host port")

	cmd.PersistentFlags().String("db.host", "", "Mongo database host")
	cmd.PersistentFlags().Int("db.port", 27017, "Mongo database port")
	cmd.PersistentFlags().String("db.database", "thesaurus", "Database name")
	cmd.PersistentFlags().String("db.login", "", "Mongo login")
	cmd.PersistentFlags().String("db.password", "", "Mongo password")

	cmd.PersistentFlags().String("csv.path", "", "Path to CVS files directory for loading into db")
	cmd.PersistentFlags().String("csv.separator.column", ",", "Separate columns in CSV")

	cmd.PersistentFlags().Bool("update", false, "Run programme in import mod")
	cmd.PersistentFlags().String("sentryDSN", "", "DNS error aggregator")

	cmd.PersistentFlags().String("logging.output", "STDOUT", "Log output target STDOUT|FILE")
	cmd.PersistentFlags().String("logging.level", "INFO", "Log lvl DEBUG|ERROR|FATAL|INFO")
	cmd.PersistentFlags().String("logging.format", "TEXT", "Log output format TEXT|JSON")

	/* Bind command flags to config variables */
	for _, val := range []string{
		"http.host",
		"http.port",
		"db.host",
		"db.port",
		"db.database",
		"db.login",
		"db.password",
		"csv.path",
		"csv.separator.column",
		"update",
		"sentryDSN",
		"logging.output",
		"logging.level",
		"logging.format",
	} {
		if err := viper.BindPFlag(val, cmd.PersistentFlags().Lookup(val)); err != nil {
			log.WithError(err).Error("Viper bind error")
		}
	}
}

func checkMandatoryParams() {
	var missing []string

	for _, val := range []string{"http.host", "db.host", "db.database"} {
		if value := viper.GetString(val); value == "" {
			missing = append(missing, val)
		}
	}

	if len(missing) != 0 {
		log.WithField("missed", missing).Fatal("Missed mandatory params. Use --help flag or config")
	}
}

func loggingSetup() {
	logLevel, err := log.ParseLevel(viper.GetString("logging.level"))
	if err != nil {
		log.WithError(err).Fatal("Can't set logging level")
	}

	log.SetLevel(logLevel)

	if viper.GetString("logging.output") != "STDOUT" {
		logFile, err := os.OpenFile(viper.GetString("logging.output"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening log file: %v", err)
		}

		log.SetOutput(logFile)

		defer func() {
			if err := logFile.Close(); err != nil {
				log.WithError(err).Error("Error while file closing ")
			}
		}()

		defer log.SetOutput(os.Stdout)
	}

	if viper.GetString("logging.format") == "JSON" {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func ravenCapture(callback func()) {
	var ravenClient *raven.Client
	var err error

	if viper.GetString("sentryDSN") != "" {
		ravenClient, err = raven.New(viper.GetString("sentryDSN"))
		if err != nil {
			log.WithError(err).Fatal("Cann't initialize sentry dns")
		}
		log.Debug("Sentry DSN: ", viper.GetString("SentryDSN"))
	} else {
		ravenClient = raven.DefaultClient
		log.Debug("Sentry DSN disable")
	}

	coreError, eventId := ravenClient.CapturePanicAndWait(func() {
		wg := &sync.WaitGroup{}
		callback()
		wg.Wait()
	}, nil)
	if eventId != "" {
		log.WithFields(log.Fields{
			"eventId": eventId,
			"error":   coreError.(error).Error(),
		}).Error("Core panic! Sending sentry event")
	}
}

func start() {
	loggingSetup()

	s := service.NewInstance(
		mongo.MustConnect(
			viper.GetString("db.host"),
			viper.GetInt("db.port"),
			viper.GetString("db.login"),
			viper.GetString("db.password"),
			viper.GetString("db.database"),
		),
	)

	ravenCapture(func() {
		log.WithFields(log.Fields{
			"context": "CORE",
			"version": version,
			"status":  "STARTED",
		}).Info("Application started successfully")

		shouldImporting := viper.GetBool("update")

		if shouldImporting {
			path := viper.GetString("csv.path")
			colSep := viper.GetString("csv.separator.column")

			csv.ImportCSV(s.Raw().(*mongo.RawMongo).Database, path, colSep)
			os.Exit(0)
		} else {
			spec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
			if err != nil {
				log.Fatal(err.Error())
			}

			server := restapi.NewServer(
				operations.NewThesaurusAPI(spec),
			)

			server.Host = viper.GetString("http.host")
			server.Port = viper.GetInt("http.port")

			server.ConfigureAPI()

			if err := server.Serve(); err != nil {
				log.Fatal(err)
			}

			server.Shutdown()
		}
	})

}

func initConfig() {
	// Use config file from the flag if provided.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Can't read config:", err)
		}
	}

	viper.SetEnvPrefix("THESAURUS")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
