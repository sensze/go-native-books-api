package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	error := viper.ReadInConfig()

	if error != nil {
		log.Fatal(error)
	}
	if error := viper.Unmarshal(&ENV); error != nil {
		log.Fatal(error)
	}

	log.Println("Config loaded")
}