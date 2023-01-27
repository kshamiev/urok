package main

import "fmt"

func main() {
	Ledger[string, int16]{
		ID:      "acct-1",
		Amounts: []int16{1, 2, 3},
		SumFn:   Sum[int16],
	}.PrintIDAndSum()

	SomeFunc[string, int](Ledger[string, int]{
		ID:      "acct-1",
		Amounts: []int{1, 2, 3},
		SumFn:   Sum[int],
	})
}

// Numeric expresses a type constraint satisfied by any numeric type.
type Numeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 |
		~int64 | ~float32 | ~float64 | ~complex64 | ~complex128
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

// Ledger is an identifiable, financial record.
type Ledger[T ~string, K Numeric] struct {

	// ID identifies the ledger.
	ID T

	// Amounts is a list of monies associated with this ledger.
	Amounts []K

	// SumFn is a function that can be used to sum the amounts
	// in this ledger.
	SumFn SumFn[K]
}

// PrintIDAndSum emits the ID of the ledger and a sum of its amounts on a
// single line to stdout.
func (l Ledger[T, K]) PrintIDAndSum() {
	fmt.Printf("%s has a sum of %v\n", l.ID, l.SumFn(l.Amounts...))
}

// ////

func SomeFunc[T ~string, K Numeric](l Ledger[T, K]) {
	l.PrintIDAndSum()
}
