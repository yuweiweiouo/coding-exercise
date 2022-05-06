package db

import (
	"os"
	"path"

	"github.com/google/wire"
	"github.com/yuweiweiouo/coding-exercise/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Provider = wire.NewSet(New)

func New(c *config.Config) (db *gorm.DB, err error) {
	os.MkdirAll(path.Dir(c.Database), os.ModePerm)
	db, err = gorm.Open(sqlite.Open(c.Database), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
