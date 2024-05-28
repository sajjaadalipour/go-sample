package usecase

import "go.uber.org/fx"

var Provide = fx.Provide(newUserService)

type UserService struct {
	userRepository UserRepository
}

func newUserService(userRepository UserRepository) UserService {
	return UserService{userRepository: userRepository}
}

func (userService UserService) Create(user User) error {
	return userService.userRepository.Save(user)
}
