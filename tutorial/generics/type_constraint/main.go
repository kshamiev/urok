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

	fmt.Printf("Generic Sums with Constraint: %v and %v and %v\n",
		SumNumberConstraint(ints),
		SumNumberConstraint(floats),
		SumNumberConstraint(floatsTest),
	)
}

// Составные ограничение типа (вынос ограничение в отдельный тип)

type NumberConstraint interface {
	int64 | float64
}

func SumNumberConstraint[K comparable, V NumberConstraint](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
