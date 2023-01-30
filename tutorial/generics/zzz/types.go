package zzz

// Numeric expresses a type constraint satisfied by any numeric type.
type Numeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~int8 | ~int16 | ~int32 |
		~int64 | ~float32 | ~float64
}

// SumFn is a type alias of a generic function
type SumFn[T Numeric] func(...T) T

type ID string

// Ledgerish expresses a constraint that may be satisfied by types that have
// ledger-like qualities.
type Ledgerish[T ~string, K Numeric] interface {
	~struct {
		ID      T
		Amounts []K
		SumFn   SumFn[K]
	}
	PrintIDAndSum()
}
