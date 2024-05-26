package main

import (
	"go.uber.org/fx"
	"test/application"
	"test/infra/config"
	"test/infra/db"
	"test/infra/repository"
	"test/infra/server"
	"test/infra/web"
)

func main() {
	app := fx.New(
		fx.Invoke(config.Init),
		fx.Provide(db.InitConnectionConfig),
		fx.Provide(db.InitGorm),
		fx.Invoke(db.RunMigration),
		// Http Server
		fx.Provide(server.InitConfig),
		fx.Provide(server.Provider),
		// User
		fx.Provide(repository.NewSQLUserRepository),
		fx.Provide(application.NewUserService),
		fx.Provide(web.NewUserController),
		// Health
		fx.Provide(web.NewHealthController),
		// Register Routes
		fx.Invoke(web.RegisterRoutes),
		// Start Server
		fx.Invoke(server.Start),
	)

	app.Run()
}
