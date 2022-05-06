//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/yuweiweiouo/coding-exercise/internal/config"
	"github.com/yuweiweiouo/coding-exercise/internal/db"
)

func CreateApp() (*App, func(), error) {
	panic(wire.Build(
		New,
		config.Provider,
		db.Provider,
	))
}
