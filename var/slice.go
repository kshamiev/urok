package main

import (
	"fmt"
)

// обьявление переменный уровня пакета
// var Srez0 []int                              // рекомендуется
var Srez1 = make([]int, 10, 15)              //
// var Srez2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // рекомендуется

func main() {
	// Srez1Func := make([]int, 10, 15)
	// fmt.Println(Srez0, Srez1Func, Srez2)

	Srez1 = append(Srez1, []int{91, 92, 93, 94, 95, 96, 97, 98, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99}...)
	Srez1[3] = 999
	fmt.Println(len(Srez1), cap(Srez1))
	sampleSlice(Srez1)

	fmt.Println(Srez1)

	sampleSliceLink(&Srez1)

	fmt.Println(Srez1)

	// Правильное копирование среза
	// var SlTest3 []int // так не правильно не скопирует
	SlTest2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	SlTest3 := make([]int, len(SlTest2), cap(SlTest2))
	copy(SlTest3, SlTest2)
}

// передача по значению
func sampleSlice(sl []int) {
	sl = append(sl, 100)
	sl[3] = 111
}

// передача по ссылке
func sampleSliceLink(sl *[]int) {
	(*sl)[3] = 555
	*sl = append(*sl, 100)
}
