package server

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func initConfig() Config {
	var config Config

	if err := viper.Sub("httpServer").Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode http-server Config into struct, %v", err)
	}

	return config
}

func newHttpServer(engine *echo.Echo, lc fx.Lifecycle) *http.Server {
	var config = initConfig()

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(config.Port),
		Handler:      engine,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := srv.ListenAndServe()
				if err != nil {
					panic(err)
				}
				fmt.Println("Starting HTTP server at", srv.Addr)
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func newEcho() *echo.Echo {
	var e = echo.New()
	e.HTTPErrorHandler = newHttpErrorHandler()
	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}

var Provide = fx.Provide(newEcho, newHttpServer)
