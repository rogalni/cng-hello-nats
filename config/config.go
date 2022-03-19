package config

import (
	"log"

	"github.com/spf13/viper"
)

var App Config

type Config struct {
	NatsUrl string `mapstructure:"NATS_URL"`
}

func init() {
	viper.SetDefault("NATS_URL", "localhost:4222")

	viper.AddConfigPath("./config")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			log.Println("Failed to load from config file")
		}
	}

	viper.AutomaticEnv()

	err := viper.Unmarshal(&App)
	if err != nil {
		panic("Could not unmarshal config")
	}
}
