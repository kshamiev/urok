// Функции
// Примеры написания функций и работы с ними
package function

import (
	"fmt"
)

type Sample func(s string) int

// вызовы примеров функций
func Function() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	fz := []int{1, 54, 87, 34}
	fmt.Println(Add(fz...))

	fmt.Println(factorial(3))

	deffer()
}

// простая функция
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// пример переменного числа аргументов одного типа (строго указывается в конце)
func Add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// Пример рекурсивной функции
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// Пример с defer
func deffer() {
	fmt.Println("Start def")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Stop def")
}

//
func closure() {
	// Функции захватывают переменные в области видимости
	// Но, чтобы передавать значение, требуется явно передавать его в функцию
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Got", i)
		}(i)
	}
	_, _ = fmt.Scanln()
}
