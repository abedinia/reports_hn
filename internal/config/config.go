package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Token string
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

	viper.AddConfigPath("/")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
