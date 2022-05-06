package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var Provider = wire.NewSet(New)

type Config struct {
	Database string `mapstructure:"database"`
}

func New() (*Config, error) {
	viper.SetConfigName("local")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	c := &Config{}
	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}

	return c, nil
}
