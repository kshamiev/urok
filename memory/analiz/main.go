package main

import (
	"time"

	"github.com/shopspring/decimal"
)

type Y struct {
	v int
	b time.Time
	m decimal.Decimal
	s string
}
type X struct{ v int }

func foo(x interface{}) {
}

// go build -gcflags=-m memory/analiz/main.go
func main() {
	x := &X{1}
	foo(x)
	y := &Y{}
	foo(y)
}
