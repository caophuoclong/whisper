package main

import (
	"fmt"

	"github.com/caophuoclong/whisper/configs"
	"github.com/caophuoclong/whisper/internal/models"
	"github.com/caophuoclong/whisper/internal/server"
	"github.com/caophuoclong/whisper/pkg/logger"
	"github.com/caophuoclong/whisper/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := configs.LoadConfig("./configs")
	if err != nil {
		fmt.Println(err)
	}
	appLogger := logger.NewLogger(cfg)
	appLogger.InitLogger()

	psqlDb, err := postgres.NewPgqlDB(cfg)
	if err != nil {
		appLogger.Fatal("Postgres init error: ", err)
		return
	}
	db, err := psqlDb.DB()

	if err != nil {
		appLogger.Error("Something wrong!!!", err)
	} else {
		appLogger.Info("Connect db successfully: ", db.Stats().InUse)
		defer db.Close()
	}
	psqlDb.AutoMigrate(&models.User{})
	gin.SetMode(gin.ReleaseMode)
	s := server.NewServer(psqlDb, appLogger, cfg)
	s.UseMiddleware(logger.GinLogger(appLogger))
	s.MapHandler()
	s.Run(cfg)
}
