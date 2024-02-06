package cmd

import (
	"dummy-server/routh"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var (
	restCmd = &cobra.Command{
		Use:   "rest",
		Short: "dummy-server",
		Run:   restServer,
	}
)

func init() {
	rootCmd.AddCommand(restCmd)
}
func restServer(cmd *cobra.Command, args []string) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{""},
		AllowHeaders: []string{""},
	}))
	routh.RouthSample(e, sampleHandler)
	routh.RouthDummyServer(e, dummyServerHandler)
	err := e.Start(rootConfig.Server.HostServer + ":" + rootConfig.Server.PortServer)
	if err != nil {
		log.Errorf("Cannot Start the application !!, Err > ", err)
		os.Exit(1)
	}

}
