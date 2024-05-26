package application

type UserRepository interface {
	Save(user User) error
}
