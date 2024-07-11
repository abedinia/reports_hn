package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"report_hn/internal/db"
)

var MigrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate tables!",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Hello from migrations")

		db.RunMigrations()
	},
}
