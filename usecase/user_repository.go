package usecase

type UserRepository interface {
	Save(user User) error
}
