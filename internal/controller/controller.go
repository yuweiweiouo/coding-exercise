package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type baseController struct {
	logger *zap.Logger
}

func (ctl baseController) success(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, gin.H{
		"result": data,
	})
}

func (ctl baseController) error(ctx *gin.Context, code int, err error) {
	ctl.logger.Error(err.Error())
	ctx.AbortWithStatusJSON(code, gin.H{
		"error": err.Error(),
	})
}
