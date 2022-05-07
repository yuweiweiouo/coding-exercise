package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BookController interface {
	All(*gin.Context)
}

func NewBookController(logger *zap.Logger) BookController {
	return &bookController{
		logger: logger,
	}
}

type bookController struct {
	logger *zap.Logger
}

func (ctl bookController) All(c *gin.Context) {
	c.JSON(200, "func (ctl bookController) All()")
}
