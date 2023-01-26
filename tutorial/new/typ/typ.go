package typ

import "github.com/google/uuid"

type Good struct {
	Test  string
	ID    int64
	Name  string
	Price float64
}

func (o *Good) Test1() {

}

type Invoice struct {
	Test    string
	ID      uuid.UUID
	Comment string
	Amount  float64
}

func (o *Invoice) Test2() {

}

type User struct {
	Test string
	ID   int64
	Fio  string
	Age  int
}

func (o *User) Test3() {

}
