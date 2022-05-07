package db

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"github.com/yuweiweiouo/coding-exercise/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Provider = wire.NewSet(New, NewOption)

type Option struct {
	Dsn string
}

func NewOption(v *viper.Viper) (*Option, error) {
	option := &Option{}

	if err := v.UnmarshalKey("database", option); err != nil {
		return nil, err
	}

	return option, nil
}

func New(option *Option) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(option.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Task{})

	return db, nil
}
