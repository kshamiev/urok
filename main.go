package main

import (
	"errors"
	"fmt"
)

type Test struct {
	Cnt int
}

func (obj Test) Add(n int) {
	obj.Cnt += n
}

func (obj Test) Del(n int) {
	obj.Cnt -= n
}

func (obj *Test) String() string {
	return fmt.Sprintf("%v", obj.Cnt)
}

func main() {
	action()
	fmt.Println("finish")
}

// action выполнение задачи
func action() {

	defer func() {
		if rvr := recover(); rvr != nil {
			fmt.Printf("panic: %+v", rvr)
		}
	}()

	err := Action()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Action() error {

	// return nil
	panic("PANICA")

	return errors.New("oshibka")

}
