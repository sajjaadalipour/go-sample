package db

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type DataSourceConfig struct {
	Host     string
	User     string
	Pass     string
	DB       string
	Port     int
	TimeZone string
	SslMode  string
}

func (c DataSourceConfig) GenerateDsn() string {

	if c.TimeZone == "" {
		c.TimeZone = `UTC`
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

func initConnectionConfig() DataSourceConfig {
	var connectionConfig DataSourceConfig

	if err := viper.Sub("db").Unmarshal(&connectionConfig); err != nil {
		fmt.Printf("Unable to decode DB Connection Config into struct, %v", err)
	}

	return connectionConfig
}

var Provide = fx.Provide(initConnectionConfig)
