package app

import (
	mysqldb "company-api/foundation/database"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App    APP
	Server Server
	DB     mysqldb.Config
}

type APP struct {
	Version     string
	Environment string
}

type Server struct {
	Port int
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../../")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
