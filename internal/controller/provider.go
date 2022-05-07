package controller

import "github.com/google/wire"

type Controllers struct {
	Task TaskController
}

var Provider = wire.NewSet(
	NewTaskController,
	wire.Struct(new(Controllers), "*"),
)
