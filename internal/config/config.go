package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var Provider = wire.NewSet(New)

type Config struct {
	Port     int
	Debug    bool
	Database string
}

func New(name string) (*Config, error) {
	viper.AddConfigPath("config")
	viper.SetConfigName(name)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	c := &Config{}
	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}

	return c, nil
}
