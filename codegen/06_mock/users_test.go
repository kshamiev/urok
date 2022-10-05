package users

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/kshamiev/urok/codegen/06_mock/mocks"
	"github.com/kshamiev/urok/codegen/06_mock/users"
)

func doSomeWork(u users.UserInterface) {
	u.SetName("Ivan Ivanov")
	name := u.GetName()
	fmt.Println(name)
	// если мы раскомментируем эту строчку, то наша последстельность нарушится и тест сфейлится
	// u.SetName("Ivan Petrov")
}

func TestDoSomethingWithUsers(t *testing.T) {
	ctrl := gomock.NewController(t) // обратите внимание - мы передаём t сюда, это надо чтобы гомок вывел корректное сообщение если тесты не пройдут
	defer ctrl.Finish()             // при завершении функции TestDoSomethingWithUsers вызовется Finish и сравнит последовательсноть вызовов

	testUser := mocks.NewMockUserInterface(ctrl)

	// тут мы записываем последовтаельность вызовов, которая должна совершиться
	testUser.EXPECT().SetName("Ivan Ivanov")
	testUser.EXPECT().GetName().Return("Ivan Ivanov")

	doSomeWork(testUser)
}
