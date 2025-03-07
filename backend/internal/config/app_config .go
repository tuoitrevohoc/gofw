package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DbUrl string `mapstructure:"db_url"`
	Port  string `mapstructure:"port"`
	Domain string `mapstructure:"domain"`
}

func LoadAppConfig() (*AppConfig, error) {
	env := "local" // default environment
	if envVar := os.Getenv("ENV"); envVar != "" {
		env = envVar
	}

	viper.SetConfigName(fmt.Sprintf("config.%s", env)) // will look for config.{env}.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../../config")

	var c AppConfig

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found: %w", err)
		}
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &c, nil
}
