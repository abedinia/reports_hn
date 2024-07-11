package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"report_hn/internal/db"
)

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed a user in user table",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Hello on seed command")

		db.SeedUser()
	},
}
