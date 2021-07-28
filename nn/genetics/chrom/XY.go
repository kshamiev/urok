package chrom

import "github.com/shopspring/decimal"

type XY struct {
	X, Y   int64
	Weight decimal.Decimal
}

type СhainXY []XY

func (g СhainXY) Len() int {
	return len(g)
}
func (g СhainXY) Less(i, j int) bool {
	return g[i].Weight.GreaterThan(g[j].Weight)
}
func (g СhainXY) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
