// Примеры управляющих операторов языка
// IF
// FOR
// SWITCH
// LABEL
// BREAK
// CONTINUE
package control

import (
	"fmt"
)

var HashControl = map[string]string{
	"qwerty1": "yuiop11",
	"qwerty2": "yuiop12",
	"qwerty3": "yuiop13",
	"qwerty4": "yuiop14",
}

func Control() {

	sampleIF()
	sampleFOR()
	sampleSWITCH()
	sampleLABEL()

}

// IF
func sampleIF() {

	if _, ok := HashControl["qwerty1"]; ok {
		fmt.Println("IF  Хеш index qwerty1 OK", "\n")
	} else {
		fmt.Println("IF  Хеш index qwerty1 NOT", "\n")
	}

}

// FOR
// главный принцип оператора, логическое выражение либо полное отсутсвие выражения
// что означет истину - или бесконечный цикл
func sampleFOR() {

	for {
		fmt.Println("FOR  Бесконечный цикл")
		break
	}

	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(sl); i++ {
		fmt.Print(sl[i], " ")
	}
	fmt.Println("\n")

	for index := range sl {
		fmt.Print(index, " ")
	}
	fmt.Println("\n")

	for index, val := range sl {
		fmt.Print(index, " = ", val, " ")
	}
	fmt.Println("\n")

}

// SWITCH
func sampleSWITCH() {

	// проверка логического условия в самом операторе
	switch HashControl["qwerty1"] {
	case "qqqq", "yuiop11":
		fmt.Println("switch qqqq, yuiop11 OK")
	case "fffff":
		if true {
			break // выходим из свича
		}
		fmt.Println("switch fffff OK")
		fallthrough // проваливаемся в следующий case
	default:
		fmt.Println("switch default OK")
	}

	// проверка логического условия в каждом case
	switch {
	case HashControl["qwerty1"] == "yuiop11":
		fmt.Println("switch case yuiop11 OK")
	case HashControl["qwerty1"] == "yuiop11": // Повторно не сработает поскольку выйдет из первого верного case
		fmt.Println("switch case yuiop11 CASE 2 OK")
	case HashControl["qwerty1"] == "fffff":
		if true {
			break // выходим из свича
		}
		fmt.Println("switch fffff OK")
		fallthrough // проваливаемся в следующий case
	default:
		fmt.Println("switch default OK")
	}
}

// LABEL
func sampleLABEL() {

MyLopp:
	for _, v := range HashControl {

		switch {
		case v == "yuiop11":
			fmt.Println("switch qqqq, yuiop11 OK")
			break MyLopp // выход будет на следующую команду после цикла FOR (то есть на следующий оператор после label)
		case v == "fffff":
			fmt.Println("switch fffff OK")
		default:
			fmt.Println("switch default OK")
		}

	}

}
