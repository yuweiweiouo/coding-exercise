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

func New(name string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName(name)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return v, nil
}
