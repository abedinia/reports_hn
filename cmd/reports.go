package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"report_hn/internal/server"
)

var ReportCmd = &cobra.Command{
	Use:   "report",
	Short: "run report api",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Hello command executed")
		server.ApiServer()
	},
}
