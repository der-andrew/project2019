package integration

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var (
	version = "No Version Provided"
	cfgFile string
)

var Cmd = &cobra.Command{
	Use:     "thesaurus",
	Short:   "Thesaurus is a service for documents management",
	Long:    `For getting more information read project wiki`,
	Version: version,
}

func Init() {
	Cmd.SetVersionTemplate(fmt.Sprintf("Version: %s\n", version))

	cobra.OnInitialize(initConfig)

	/* Set command flags */
	Cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")

	Cmd.PersistentFlags().String("http.host", "0.0.0.0", "Host address")
	Cmd.PersistentFlags().Int("http.port", 80, "Host port")

	Cmd.PersistentFlags().String("db.host", "", "Mongo database host")
	Cmd.PersistentFlags().Int("db.port", 27017, "Mongo database port")
	Cmd.PersistentFlags().String("db.database", "thesaurus", "Database name")
	Cmd.PersistentFlags().String("db.login", "", "Mongo login")
	Cmd.PersistentFlags().String("db.password", "", "Mongo password")

	Cmd.PersistentFlags().String("import", "", "Path to CVS files directory for loading into db")
	Cmd.PersistentFlags().String("sentryDSN", "", "DNS error aggregator")

	Cmd.PersistentFlags().String("logging.output", "STDOUT", "Log output target STDOUT|FILE")
	Cmd.PersistentFlags().String("logging.level", "INFO", "Log lvl DEBUG|ERROR|FATAL|INFO")
	Cmd.PersistentFlags().String("logging.format", "TEXT", "Log output format TEXT|JSON")

	/* Bind command flags to config variables */
	for _, val := range []string{
		"http.host",
		"http.port",
		"db.host",
		"db.port",
		"db.database",
		"db.login",
		"db.password",
		"documentsPath",
		"sentryDSN",
		"logging.output",
		"logging.level",
		"logging.format",
	} {
		if err := viper.BindPFlag(val, Cmd.PersistentFlags().Lookup(val)); err != nil {
			log.WithError(err).Error("Viper bind error")
		}
	}
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
