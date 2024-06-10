package gorm

import (
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"test/db"
)

func initGorm(config db.DataSourceConfig) (*gorm.DB, error) {
	dsn := config.GenerateDsn()
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

var Provide = fx.Provide(initGorm)
