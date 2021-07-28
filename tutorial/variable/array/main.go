package main

import (
	"fmt"
)

var arr0 [4]int
var arr1 = [5]int{3, 5, 6, 7, 9}
var arr2 = [...]int{3, 5, 6, 7, 9} // вместо точек можно написать реальный размер массива

func main() {
	arr0[2] = 999

	fmt.Println(arr0, arr1, arr2)

	sampleArr(arr1)
	fmt.Println(arr1)

	sampleArrLink(&arr1)
	fmt.Println(arr1)

	// получение или создание среза от массива
	ar := [...]int{5, 6, 7}
	sl := ar[:] // получится уже срез
	sl = append(sl, 3)
	fmt.Println("получение или создание среза от массива ", ar, sl)

}

// передача по значению
func sampleArr(arr [5]int) {
	arr[2] = 100
}

// передача по ссылке
func sampleArrLink(arr *[5]int) {
	arr[2] = 100
}
