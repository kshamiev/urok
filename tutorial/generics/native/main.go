package main

import (
	"fmt"
)

type Test struct {
	Name string
}

func main() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	floats := map[int]float64{
		1: 35.98,
		2: 26.99,
	}
	floatsTest := map[Test]float64{
		Test{}: 35.98,
		Test{}: 26.99,
	}

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v and %v\n",
		SumNumber(ints),
		SumNumber(floats),
		SumNumber(floatsTest),
	)
}

// Простой вариант (объявление ограничения типа в самой функции)

func SumNumber[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
