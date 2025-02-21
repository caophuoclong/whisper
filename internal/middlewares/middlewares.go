package middlewares

import (
	"github.com/caophuoclong/whisper/configs"
	"github.com/caophuoclong/whisper/internal/auth"
	"github.com/caophuoclong/whisper/pkg"
)

type MiddlewareManager struct {
	logger pkg.Logger
	au     auth.AuthUsecase
	cfg    *configs.Config
}

func NewMiddlewareManager(l pkg.Logger, au auth.AuthUsecase, cfg *configs.Config) MiddlewareManager {
	return MiddlewareManager{
		logger: l,
		au:     au,
		cfg:    cfg,
	}
}
