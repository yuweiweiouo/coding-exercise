package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"gorm.io/gorm"
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
	option *Option
	router *gin.Engine
	Db     *gorm.DB
}

func (s Server) Start() error {
	return s.router.Run(fmt.Sprintf(":%d", s.option.Port))
}

func New(db *gorm.DB, option *Option, router *gin.Engine) *Server {
	return &Server{
		option: option,
		router: router,
		Db:     db,
	}
}
