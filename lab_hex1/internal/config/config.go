package config

import (
	"log"

	"github.com/spf13/viper"
)

type SQLServerConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type CacheConfig struct {
	RefreshMinutes int
}

type ExternalAPIConfig struct {
	URL string
}

type Config struct {
	SQLServer   SQLServerConfig
	Cache       CacheConfig
	ExternalAPI ExternalAPIConfig
}

func LoadConfig() *Config {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return &config
}
