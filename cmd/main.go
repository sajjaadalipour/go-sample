package main

import (
	"go.uber.org/fx"
	"net/http"
	"test/config"
	"test/db"
	"test/external/gorm"
	"test/repository"
	"test/rest"
	"test/server"
	"test/usecase"
)

func main() {
	app := fx.New(
		db.Provide,
		gorm.Provide,
		server.Provide,
		repository.Provide,
		usecase.Provide,
		rest.Provide,
		fx.Invoke(config.Init, db.RunMigration, rest.RegisterRoutes, func(*http.Server) {}),
	)

	app.Run()
}
