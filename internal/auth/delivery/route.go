package delivery

import (
	"github.com/caophuoclong/whisper/internal/auth"
	"github.com/caophuoclong/whisper/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func MapAuthRouter(authGroup *gin.RouterGroup, authHandler auth.Handlers, mw middlewares.MiddlewareManager) {

	authGroup.GET("/health", authHandler.Health())
	authGroup.POST("/register",
		authHandler.Register())
	authGroup.POST("/login", authHandler.Login())
	authGroup.Use(mw.AuthJWTFromRequest())
	authGroup.GET("/me", authHandler.GetMe())
}
