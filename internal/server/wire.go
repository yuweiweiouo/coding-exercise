//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/yuweiweiouo/coding-exercise/internal/config"
	"github.com/yuweiweiouo/coding-exercise/internal/controller"
	"github.com/yuweiweiouo/coding-exercise/internal/dao"
	"github.com/yuweiweiouo/coding-exercise/internal/db"
	"github.com/yuweiweiouo/coding-exercise/internal/log"
	"github.com/yuweiweiouo/coding-exercise/internal/router"
	"github.com/yuweiweiouo/coding-exercise/internal/service"
)

func CreateServer(configName string) (*Server, func(), error) {
	panic(wire.Build(
		Provider,
		config.Provider,
		db.Provider,
		controller.Provider,
		router.Provider,
		log.Provider,
		service.Provider,
		dao.Provider,
	))
}
