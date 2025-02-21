package server

import (
	"github.com/caophuoclong/whisper/internal/auth/delivery"
	"github.com/caophuoclong/whisper/internal/auth/repository"
	"github.com/caophuoclong/whisper/internal/auth/usecase"
	"github.com/caophuoclong/whisper/internal/middlewares"
)

func (s *Server) MapHandler() {

	v1Group := s.s.Group("/api/v1")

	// Router
	authGroup := v1Group.Group("/auth")

	//Repo
	authRepo := repository.NewAuthRepo(s.db)

	//Usecase
	authUsecase := usecase.NewAuthUsecase(authRepo, s.cfg)

	mw := middlewares.NewMiddlewareManager(
		s.logger,
		authUsecase,
		s.cfg,
	)
	//Handler
	authHandler := delivery.NewAuthHandler(authUsecase, s.logger)
	delivery.MapAuthRouter(authGroup, authHandler, mw)
}
