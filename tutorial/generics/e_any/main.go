package main

import "fmt"

// any
// это синоним interface{}

func main() {
	//
	fmt.Println(
		Ptr("local"), *Ptr("local"),
	)
	fmt.Println(
		Ptr(80), *Ptr(80),
	)

	//
	var v Vector[int]
	v.Push(5)
	v.Push(55)
	v.Push(555)
	fmt.Println(v)
}

// Ptr returns *value.
func Ptr[T any](value T) *T {
	return &value
}

// ////

type Vector[T any] []T

func (v *Vector[T]) Push(x T) {
	*v = append(*v, x)
}
