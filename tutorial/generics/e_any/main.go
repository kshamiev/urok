package main

import "fmt"

// any
// это синоним interface{}

func main() {
	fmt.Println(
		Ptr("local"), *Ptr("local"),
	)
	fmt.Println(
		Ptr(80), *Ptr(80),
	)
}

// Ptr returns *value.
func Ptr[T any](value T) *T {
	return &value
}
