package zzz

import "fmt"

// Sum returns the sum of the provided arguments.
func Sum[T Numeric](args ...T) T {
	var sum T
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

// Calc is an identifiable, financial record.
type Calc[T ~string, K Numeric] struct {

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
func (l Calc[T, K]) PrintIDAndSum() {
	fmt.Printf("%s has a sum of %v\n", l.ID, l.SumFn(l.Amounts...))
}

// PrintLedger emits a ledger's ID and total amount on a single line
// to stdout.
func PrintLedger[T ~string, K Numeric, L Ledgerish[T, K]](l L) {
	l.PrintIDAndSum()
}
