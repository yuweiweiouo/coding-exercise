package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Provider = wire.NewSet(New, NewOption)

type Option struct {
	Port int
}

func NewOption(v *viper.Viper) (*Option, error) {
	var (
		err    error
		option = &Option{}
	)

	if err = v.UnmarshalKey("server", option); err != nil {
		return nil, err
	}

	return option, err
}

type Server struct {
	logger *zap.Logger
	option *Option
	router *gin.Engine
}

func (s Server) Start() error {
	addr := fmt.Sprintf(":%d", s.option.Port)
	s.logger.Info("伺服器啟動 => " + addr)
	return s.router.Run(addr)
}

func New(logger *zap.Logger, option *Option, router *gin.Engine) *Server {
	return &Server{
		option: option,
		router: router,
		logger: logger,
	}
}
