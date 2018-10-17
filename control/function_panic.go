// Обработка паники
package control

import (
	"errors"
	"fmt"
)

func PanicTest() (err error) {

	defer recoveryTest(&err)

	fmt.Println("Start Test Panic")
	panic("Шеф все пропало !")
	return
}

func recoveryTest(err *error) {

	if e := recover(); e != nil {
		*err = errors.New(fmt.Sprintf("%v", e))
	}
}
