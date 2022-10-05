package users

// если мы хотим тестить юзера через моки, то мы должны положить тесты в отдлельный файл

type User struct {
	ID   int
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(newName string) {
	u.Name = newName
}

//go:generate mockgen -package=mocks -destination=../mocks/users_mock.go -source=./users.go UserInterface
type UserInterface interface {
	GetName() string
	SetName(string)
}
