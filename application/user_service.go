package application

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return UserService{userRepository: userRepository}
}

func (userService UserService) Create(user User) error {
	return userService.userRepository.Save(user)
}
