package config

import (
	"go-app/pkg/logger"

	"github.com/spf13/viper"
)

// SlackConfig to hold serect/key/config of slack
type SlackConfig struct {
	SigningSecret string `mapstructure:"SLACK_SIGNING_SECRET"`
	SlackToken    string `mapstructure:"SLACK_TOKEN"`
	SlackAPI      string `mapstructure:"SLACK_SEND_MESSAGE_API"`
}

// GetSlackConfig Unmarshal SlackConfig from env
func GetSlackConfig() SlackConfig {
	c := SlackConfig{}
	if err := viper.Unmarshal(&c); err != nil {
		logger.Error().Fatal(err)
	}

	return c
}
