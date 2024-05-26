package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionConfig struct {
	Host     string
	User     string
	Pass     string
	DB       string
	Port     int
	TimeZone string
	SslMode  string
}

func (c ConnectionConfig) GenerateDsn() string {
	if c.Port == 0 {
		c.Port = 5432
	}

	if c.TimeZone == "" {
		c.TimeZone = `Asia/Shanghai`
	}

	if c.SslMode == "" {
		c.SslMode = `disable`
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.Host,
		c.User,
		c.Pass,
		c.DB,
		c.Port,
		c.SslMode,
		c.TimeZone,
	)
}

func (c ConnectionConfig) GenerateDataSourceName() string {
	if c.Port == 0 {
		c.Port = 5432
	}

	if c.TimeZone == "" {
		c.TimeZone = `Asia/Shanghai`
	}

	if c.SslMode == "" {
		c.SslMode = `disable`
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&timezone=%s",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.DB,
		c.SslMode,
		c.TimeZone,
	)
}

func InitConnectionConfig() ConnectionConfig {
	var connectionConfig ConnectionConfig

	if err := viper.Sub("db").Unmarshal(&connectionConfig); err != nil {
		fmt.Printf("Unable to decode DB Connection Config into struct, %v", err)
	}

	return connectionConfig
}

func InitGorm(config ConnectionConfig) (*gorm.DB, error) {
	dsn := config.GenerateDsn()
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
