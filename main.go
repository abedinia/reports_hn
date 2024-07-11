package main

import (
	"fmt"
	"os"
	"report_hn/internal/config"
	"report_hn/internal/logger"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"report_hn/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "report of hn",
	Short: "get reports of customers about dataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running My Application")
	},
}

func initConfigLogs() {
	config.LoadConfig()
	logger.InitLogger()
	logrus.Info("Configuration loaded successfully")
}

func init() {
	cobra.OnInitialize(initConfigLogs)
	rootCmd.AddCommand(cmd.MigrationCmd)
	rootCmd.AddCommand(cmd.SeedCmd)
	rootCmd.AddCommand(cmd.ReportCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Log.Errorf("Error executing root command: %v", err)
		os.Exit(1)
	}
}
