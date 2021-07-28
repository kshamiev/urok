// Обработка паники
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Start Test Panic")
	fmt.Println(PanicTest())
	fmt.Println("OK")
}

func PanicTest() (err error) {
	defer recoveryTest(&err)
	panic("Шеф все пропало !")
	// сюда никогда не дойдем
	return err
}

func recoveryTest(err *error) {
	if e := recover(); e != nil {
		*err = errors.New(fmt.Sprintf("%v", e))
	}
}
