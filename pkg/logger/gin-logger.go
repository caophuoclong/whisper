package logger

import (
	"time"

	"github.com/caophuoclong/whisper/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GinLogger(l pkg.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		l.Info(
			"Request",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.Int("status", ctx.Writer.Status()),
			zap.String("client_ip", ctx.ClientIP()),
			zap.Duration("latency", time.Since(start)),
		)
	}

}
