package config

import (
	"go-app/pkg/logger"

	"github.com/spf13/viper"
)

// SlackConfig to hold serect/key/config of slack
type SlackConfig struct {
	SigningSecret string `mapstructure:"SLACK_SIGNING_SECRET"`
}

// GetSlackConfig Unmarshal SlackConfig from env
func GetSlackConfig() SlackConfig {
	c := SlackConfig{}
	if err := viper.Unmarshal(&c); err != nil {
		logger.Error().Fatal(err)
	}

	return c
}
