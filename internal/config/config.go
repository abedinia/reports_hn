package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Env  string
	}
	Log struct {
		Level  string
		Format string
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Dbname   string
	}
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
