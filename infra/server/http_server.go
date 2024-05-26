package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func InitConfig() Config {
	var config Config

	if err := viper.Sub("http_server").Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode http-server Config into struct, %v", err)
	}

	return config
}

func Provider() *echo.Echo {
	router := echo.New()

	router.HTTPErrorHandler = GlobalHttpErrorHandler()

	return router
}

func Start(config Config, engine *echo.Echo) {
	s := &http.Server{
		Addr:         ":" + strconv.Itoa(config.Port),
		Handler:      engine,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
