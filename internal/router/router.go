package router

import (
	"time"

	"github.com/gin-contrib/pprof"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"github.com/yuweiweiouo/coding-exercise/internal/controller"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(New, NewOption)

type Option struct {
	Pprof bool
	Mode  string
}

func NewOption(v *viper.Viper) (*Option, error) {
	option := &Option{
		Mode: gin.ReleaseMode,
	}

	if err := v.UnmarshalKey("router", option); err != nil {
		return nil, err
	}

	return option, nil
}

type AddRoute func(r *gin.Engine)

func New(option *Option, ctls *controller.Controllers, logger *zap.Logger) *gin.Engine {
	if option.Mode != "" {
		gin.SetMode(option.Mode)
	}

	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	if option.Pprof {
		pprof.Register(r)
	}

	addBookRoute(r, ctls.Book)

	return r
}

func addBookRoute(r *gin.Engine, ctl controller.BookController) {
	r.GET("book", ctl.All)
}
