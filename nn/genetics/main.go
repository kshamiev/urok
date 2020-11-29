package main

import (
	"fmt"
	"sort"

	"github.com/shopspring/decimal"
)

type Gen struct {
	X, Y int64
	Z    decimal.Decimal
}

type Gens []Gen

func (g Gens) Len() int {
	return len(g)
}
func (g Gens) Less(i, j int) bool {
	return g[i].Z.GreaterThan(g[j].Z)
}
func (g Gens) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

type Chromosome struct {
}

func main() {
	genom := Gens{
		{X: -2, Y: 0},
		{X: -1, Y: -2},
		{X: 0, Y: -1},
		{X: 2, Y: 1},
	}

	for i := range genom {
		genom[i].Z = formulaExt(genom[i].X, genom[i].Y)
	}

	for i := range genom {
		fmt.Println(genom[i].Z)
	}

	fmt.Println()

	sort.Sort(genom)

	for i := range genom {
		fmt.Println(genom[i].Z)
	}

}

// Z = x / x^2+2*y^2+1
func formulaExt(x, y int64) decimal.Decimal {
	x1 := decimal.NewFromInt(x)
	y1 := decimal.NewFromInt(y)
	two := decimal.NewFromInt(2)
	return x1.Div(x1.Pow(two).Add(y1.Pow(two).Mul(two)).Add(decimal.NewFromInt(1)))
}
