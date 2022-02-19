// Примеры управляющих операторов языка
// IF		логика
// SWITCH	логика
// FOR		циклы + логика
// LABEL	goto
// BREAK	выход из блока
// CONTINUE	пропуск итерации блока
package main

import (
	"fmt"
)

func main() {
	obj := Instant{}
	sampleIF(obj)
}

// IF
func sampleIF(obj Facer) {
	if _, ok := HashControl["qwerty1"]; ok {
		fmt.Println("IF  Хеш index qwerty1 OK")
	} else {
		fmt.Println("IF  Хеш index qwerty1 NOT")
	}

	if elm, ok := obj.(Instant); ok {
		elm.Load()
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
		// 1 инициализация индекса
		// 2 проверка условия
		// 3 итерация
		// 4 увеличение индекса
		fmt.Print(sl[i], " ")
	}
	fmt.Println()

	for index := range sl {
		fmt.Print(index, " ")
	}
	fmt.Println()

	for index, val := range sl {
		fmt.Print(index, " = ", val, " ")
	}
	fmt.Println()

}

// SWITCH
func sampleSWITCH(obj Facer) {

	switch t := obj.(type) {
	default:
		fmt.Println("default type assertion")
	case Instant:
		t.Load()
	}

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
	for _, v := range HashControl {
		switch {
		case v == "yuiop11":
			fmt.Println("switch yuiop11 OK")
			goto MyLopp // выход будет на следующую команду после цикла FOR (то есть на следующий оператор после label)
		case v == "fffff":
			fmt.Println("switch fffff OK")
		default:
			fmt.Println("switch default OK")
		}
	}
	fmt.Println("GOTO 1")
MyLopp: // переходим сюда
	fmt.Println("GOTO 2")
}

// ////

type Instant struct {
}

func (Instant) Load() {
	fmt.Println("implement method LOAD")
}

type Facer interface {
	Load()
}

var HashControl = map[string]string{
	"qwerty1": "yuiop11",
	"qwerty2": "yuiop12",
	"qwerty3": "yuiop13",
	"qwerty4": "yuiop14",
}
