package main

import (
	"fmt"
	"sort"

	"github.com/kshamiev/urok/nn/genetics/chrom"
	"github.com/kshamiev/urok/nn/genetics/formula"
)

type Evolution struct {
	chain chrom.СhainXY
}

func (e *Evolution) Selection() {
	for i := range e.chain {
		e.chain[i].Weight = formula.Max(e.chain[i].X, e.chain[i].Y)
	}
	sort.Sort(e.chain)
	fmt.Println(e.chain)
}
func main() {
	test := &Evolution{chain: chrom.СhainXY{
		{X: -2, Y: 0},
		{X: -1, Y: -2},
		{X: 0, Y: -1},
		{X: 2, Y: 1},
	}}
	test.Selection()
}
