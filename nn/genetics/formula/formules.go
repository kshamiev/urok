package formula

import "github.com/shopspring/decimal"

// Z = x / x^2+2*y^2+1
func Max(x, y int64) decimal.Decimal {
	x1 := decimal.NewFromInt(x)
	y1 := decimal.NewFromInt(y)
	two := decimal.NewFromInt(2)
	return x1.Div(x1.Pow(two).Add(y1.Pow(two).Mul(two)).Add(decimal.NewFromInt(1)))
}
