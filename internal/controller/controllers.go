package controller

import "github.com/google/wire"

type Controllers struct {
	Book BookController
}

var Provider = wire.NewSet(
	NewBookController,
	wire.Struct(new(Controllers), "*"),
)
