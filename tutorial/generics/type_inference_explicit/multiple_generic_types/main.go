package main

import (
	"fmt"
)

// Numeric expresses a type constraint satisfied by any numeric type.
type Numeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 |
		~complex64 | ~complex128
}

// Sum returns the sum of the provided arguments.
func Sum[T Numeric](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// SumFn is a type alias of a generic function
type SumFn[T Numeric] func(...T) T

// id is a type alias for a string
type id string

// PrintIDAndSum prints the provided ID and sum of the given values to stdout.
func PrintIDAndSum[T ~string, K Numeric](id T, sum SumFn[K], values ...K) {

	// The format string uses "%v" to emit the sum since using "%d" would
	// be invalid if the value type was a float or complex variant.
	fmt.Printf("%s has a sum of %v\n", id, sum(values...))
}

func main() {
	PrintIDAndSum("acct-1", Sum[int], 1, 2, 3)
}
