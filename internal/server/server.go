package server

import (
	"github.com/caophuoclong/whisper/configs"
	"github.com/caophuoclong/whisper/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	s      *gin.Engine
	db     *gorm.DB
	logger pkg.Logger
	cfg    *configs.Config
}

func NewServer(db *gorm.DB, logger pkg.Logger, cfg *configs.Config) *Server {
	return &Server{
		s:      gin.Default(),
		db:     db,
		logger: logger,
		cfg:    cfg,
	}
}

func (s *Server) Run(cfg *configs.Config) {
	s.s.Run(
		cfg.Http.Port,
	)
}

func (s *Server) UseMiddleware(mdw gin.HandlerFunc) {
	s.s.Use(mdw)
}
