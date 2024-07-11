package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	logLevel, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		logrus.Fatalf("Error parsing log level, %s", err)
	}

	Log.SetLevel(logLevel)

	switch viper.GetString("log.format") {
	case "json":
		Log.SetFormatter(&logrus.JSONFormatter{})
	default:
		Log.SetFormatter(&logrus.TextFormatter{})
	}

	Log.Info("Logger initialized successfully")
}
