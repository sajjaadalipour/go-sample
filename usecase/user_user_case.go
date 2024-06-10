package usecase

import (
	"go.uber.org/fx"
	"test/repository/user"
)

var Provide = fx.Provide(newUserService)

type UserUseCase struct {
	userRepository repository.UserRepository
}

func newUserService(userRepository repository.UserRepository) UserUseCase {
	return UserUseCase{userRepository: userRepository}
}

func (userUseCase UserUseCase) Create(user repository.User) error {
	return userUseCase.userRepository.Save(user)
}
