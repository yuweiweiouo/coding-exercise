//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/yuweiweiouo/coding-exercise/internal/dao"
)

func CreateTaskService(taskDao dao.TaskDao) (TaskService, error) {
	panic(wire.Build(
		Provider,
	))
}
