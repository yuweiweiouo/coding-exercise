package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/yuweiweiouo/coding-exercise/internal/config"
	"github.com/yuweiweiouo/coding-exercise/internal/controller"
)

type AddRoute func(r *gin.Engine)

var Provider = wire.NewSet(New)

func New(c *config.Config, ctls *controller.Controllers) *gin.Engine {
	if c.Debug {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	if c.Debug {
		pprof.Register(r)
	}

	addBookRoute(r, ctls.Book)

	return r
}

func addBookRoute(r *gin.Engine, ctl controller.BookController) {
	r.GET("book", ctl.All)
}
