package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yuweiweiouo/coding-exercise/internal/config"
	"gorm.io/gorm"
)

type Server struct {
	port   int
	router *gin.Engine
	Db     *gorm.DB
}

func (s Server) Start() error {
	return s.router.Run(fmt.Sprintf(":%d", s.port))
}

func New(db *gorm.DB, conf *config.Config, router *gin.Engine) *Server {
	return &Server{
		port:   conf.Port,
		router: router,
		Db:     db,
	}
}
