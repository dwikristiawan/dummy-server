package cmd

import (
	config "dummy-server/connfig"
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
	rootConfig    *config.Root
	database      *sqlx.DB
	sampleHandler sample.Handler
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
		initApp()
	})
}
func configReader() {
	log.Infof("Initialize ENV")
	rootConfig = config.Load(EnvFilePath)
}
func initApp() {
	initSample()

}
func initSample() {
	log.Infof("Initialize sample module")
	repo := sample.NewRepository(database)
	service := sample.NewService(repo)
	controller := sample.NewController(service)
	sampleHandler = sample.NewHandler(controller)

}
