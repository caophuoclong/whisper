package postgres

import (
	"fmt"

	"github.com/caophuoclong/whisper/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewPgqlDB(config *configs.Config) (*gorm.DB, error) {
	datasourceName := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.Name,
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.SSL,
	)
	db, err := gorm.Open(postgres.Open(datasourceName), &gorm.Config{})

	sqlDb, er := db.DB()
	if er == nil {
		sqlDb.SetConnMaxIdleTime(maxIdleConns)
		sqlDb.SetMaxOpenConns(maxOpenConns)
		sqlDb.SetMaxIdleConns(maxIdleConns)
		sqlDb.SetConnMaxLifetime(connMaxLifetime)
	}
	return db, err
}
