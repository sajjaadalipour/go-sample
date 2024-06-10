package repository

import (
	"go.uber.org/fx"
	repository "test/repository/user"
)

var Provide = fx.Provide(repository.NewUserRepository)
