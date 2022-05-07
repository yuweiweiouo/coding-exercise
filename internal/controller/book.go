package controller

import (
	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(*gin.Context)
}

func NewBookController() BookController {
	return &bookController{}
}

type bookController struct {
}

func (ctl bookController) All(c *gin.Context) {
	c.JSON(200, "func (ctl bookController) All()")
}
