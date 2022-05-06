package app

import "gorm.io/gorm"

type App struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *App {
	return &App{
		Db: db,
	}
}
