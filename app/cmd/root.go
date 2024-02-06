package cmd

import (
	config "dummy-server/connfig"
	"dummy-server/module/dummyServer"
	"dummy-server/module/sample"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var (
	EnvFilePath string
	rootCmd     = &cobra.Command{
		Use:   "cobra-cli",
		Short: "dummy-server",
	}
)
var (
	rootConfig         *config.Root
	database           *sqlx.DB
	sampleHandler      sample.Handler
	dummyServerHandler dummyServer.Handler
)

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&EnvFilePath, "env", "e", ".env", ".env file to read from")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Cannot Run CLI. err > ", err)
		os.Exit(1)
	}
}
func init() {
	cobra.OnInitialize(func() {
		configReader()
		initPostgres()
		initApp()
	})
}
func configReader() {
	log.Infof("Initialize ENV")
	rootConfig = config.Load(EnvFilePath)
}
func initApp() {
	initSample()
	initDummyServer()

}
func initSample() {
	log.Infof("Initialize sample module")
	repo := sample.NewRepository(database)
	service := sample.NewService(repo)
	controller := sample.NewController(service)
	sampleHandler = sample.NewHandler(controller)

}
func initPostgres() {
	log.Infof("Initialize postgress")
	var err error
	database, err = config.OpenPostgresDatabaseConnection(config.Postgres{
		Host:                  rootConfig.Postgres.Host,
		Port:                  rootConfig.Postgres.Port,
		User:                  rootConfig.Postgres.User,
		Password:              rootConfig.Postgres.Password,
		Dbname:                rootConfig.Postgres.Dbname,
		MaxConnectionLifetime: rootConfig.Postgres.MaxConnectionLifetime,
		MaxOpenConnection:     rootConfig.Postgres.MaxOpenConnection,
		MaxIdleConnection:     rootConfig.Postgres.MaxIdleConnection,
	})
	if err != nil {
		log.Errorf("Posgress failed, error: ", err)
	}
}

func initDummyServer() {
	log.Infof("Initialize dammyServer module")
	repo := dummyServer.NewRepository(database)
	service := dummyServer.NewService(repo)
	controller := dummyServer.NewController(service)
	dummyServerHandler = dummyServer.NewHandler(controller)
}
