package main

import (
	"fmt"

	"github.com/kshamiev/urok/tutorial/generics/zzz"
)

func main() {
	Ledger[string, int16]{
		ID:      "acct-1",
		Amounts: []int16{1, 2, 3},
		SumFn:   zzz.Sum[int16],
	}.PrintIDAndSum()

	SomeFunc[string, int](Ledger[string, int]{
		ID:      "acct-1",
		Amounts: []int{1, 2, 3},
		SumFn:   zzz.Sum[int],
	})
}

// Ledger is an identifiable, financial record.
type Ledger[T ~string, K zzz.Numeric] struct {

	// ID identifies the ledger.
	ID T

	// Amounts is a list of monies associated with this ledger.
	Amounts []K

	// SumFn is a function that can be used to sum the amounts
	// in this ledger.
	SumFn zzz.SumFn[K]
}

// PrintIDAndSum emits the ID of the ledger and a sum of its amounts on a
// single line to stdout.
func (l Ledger[T, K]) PrintIDAndSum() {
	fmt.Printf("%s has a sum of %v\n", l.ID, l.SumFn(l.Amounts...))
}

// ////

func SomeFunc[T ~string, K zzz.Numeric](l Ledger[T, K]) {
	l.PrintIDAndSum()
}
