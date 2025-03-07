package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Url string `mapstructure:"url"`
}

type AppConfig struct {
	Db   DbConfig `mapstructure:"db"`
	Port string   `mapstructure:"port"`
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

	viper.AutomaticEnv()

	var c AppConfig

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found: %w", err)
		}
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &c, nil
}
