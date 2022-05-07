// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/yuweiweiouo/coding-exercise/internal/config"
	"github.com/yuweiweiouo/coding-exercise/internal/controller"
	"github.com/yuweiweiouo/coding-exercise/internal/dao"
	"github.com/yuweiweiouo/coding-exercise/internal/db"
	"github.com/yuweiweiouo/coding-exercise/internal/mylog"
	"github.com/yuweiweiouo/coding-exercise/internal/router"
	"github.com/yuweiweiouo/coding-exercise/internal/service"
)

// Injectors from wire.go:

func CreateServer(configName string) (*Server, func(), error) {
	viper, err := config.New(configName)
	if err != nil {
		return nil, nil, err
	}
	option, err := mylog.NewOption(viper)
	if err != nil {
		return nil, nil, err
	}
	logger, err := mylog.New(option)
	if err != nil {
		return nil, nil, err
	}
	serverOption, err := NewOption(viper)
	if err != nil {
		return nil, nil, err
	}
	routerOption, err := router.NewOption(viper)
	if err != nil {
		return nil, nil, err
	}
	dbOption, err := db.NewOption(viper)
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := db.New(dbOption)
	if err != nil {
		return nil, nil, err
	}
	taskDao := dao.NewTaskDao(gormDB, logger)
	taskService := service.NewTaskService(taskDao)
	taskController := controller.NewTaskController(logger, taskService)
	controllers := &controller.Controllers{
		Task: taskController,
	}
	engine := router.New(routerOption, controllers, logger)
	server := New(logger, serverOption, engine)
	return server, func() {
	}, nil
}
