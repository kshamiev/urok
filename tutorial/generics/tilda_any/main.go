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
	StatusTest := map[Test]Status{
		Test{}: Status(35),
		Test{}: Status(26),
	}

	fmt.Printf("Generic Sums with Constraint: %v and %v and %v\n",
		SumNumberConstraint(ints),
		SumNumberConstraint(floats),
		SumNumberConstraint(StatusTest),
	)

	fmt.Println(
		Ptr("local"), *Ptr("local"),
		Ptr(80), *Ptr(80),
	)
}

// Ограничение типа (вынос ограничение в отдельный тип использования в других функциях)

type Status int

type NumberConstraint interface {
	int64 | float64 | ~int
}

func SumNumberConstraint[K comparable, V NumberConstraint](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Ptr returns *value.
func Ptr[T any](value T) *T {
	return &value
}
