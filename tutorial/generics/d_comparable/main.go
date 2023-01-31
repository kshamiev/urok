package main

import "fmt"

type Test struct {
	Name string
}

// comparable
// это интерфейс-ограничение сопоставимых базовых типов и структур

func main() {
	// Математика
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[int]float64{
		1: 35.98,
		2: 26.99,
	}
	test := map[Test]float64{
		Test{}: 35.98,
		Test{}: 26.99,
	}
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v and %v\n",
		SumNumber(ints),
		SumNumber(floats),
		SumNumber(test),
	)

	// Сравнения
	aa := []int{
		2, 4, 6,
	}
	fmt.Println(Index(aa, 4))
	fmt.Println(Index(aa, 5))

	aaa := []int{
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	}
	Reverse(aaa)
	fmt.Println(aaa)
}

func SumNumber[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func Reverse[T comparable](s []T) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

// ВАЖНО !
// Все перечисляемые ограничения типа конкретного дженерика должны поддерживать общие операции
// Или над дженериками можно производить операции которые поддерживаются всеми типами указанными в ограничении

func Test1[T int | int16 | int32 | int64](a, b T) T {
	return a + b
}

func Test2[T comparable](a, b T) bool {
	return a == b
}
