package raszdel1

import (
	"math"

	"github.com/shopspring/decimal"
)

// 100 уроков математики от Алексея Савватеева!

// 1
// Теорема пифагора a^2+b^2=c^2

func Pythagoras(a, b int) (c float64) {
	return math.Sqrt(float64(a*a + b*b))
}

// 2
// Геометрический алгоритм Евклида
// Соизмеримость и не соизмеримость отрезков

func Euclid(a, b float64) bool {
	var d1, d2, d3 decimal.Decimal
	if a >= b {
		d1 = decimal.NewFromFloat(a)
		d2 = decimal.NewFromFloat(b)
	} else {
		d1 = decimal.NewFromFloat(b)
		d2 = decimal.NewFromFloat(a)
	}
	for {
		d3 = d1.Mod(d2)
		// fmt.Println("d3 остаток от деления:", d1, d2, d3, d1.Div(d2))
		if d3.Equal(decimal.NewFromInt(0)) {
			return true
		}

		n1 := d1.Div(d2).Truncate(9)
		n1 = n1.Sub(n1.Floor())
		n2 := d2.Div(d3).Truncate(9)
		n2 = n2.Sub(n2.Floor())
		if n1.Equal(n2) {
			return false // пары чисел бесконечно пропорциональны
		}

		d1 = d2
		d2 = d3
	}
}

// 3
// Сумма последовательных чисел по Гаусу
// n*(n+1)/2

func Gaus(n int) int {
	return n * (n + 1) / 2
}

// Сумма подряд идущих нечётных чисел всегда является квадратом их количества
// n(count) = n^2
// 1 = 1 = 1
// 3 = 2 = 4
// 5 = 3 = 9
// 7 = 4 = 16
// 9 = 5 = 25
// 11 = 6 = 36

func SquareOfTheNumber(nList ...int) float64 {
	var n int
	for _, v := range nList {
		n += v
	}
	return math.Sqrt(float64(n))
}

// Распределительный закон
// a^2 - b^2 = (a-b)(a+b)

// Беном Ньютона
// (a+b)^2 = a^2 + 2ab + b^2
// (a+b)^3 = a^3 + 3a^2b + 3ab^2 + b^3
// (a+b)^4 = a^4 + 4a^3b + 6a^2b^2 + 4ab^3 + b^4

func Benom(a, b float64, n int) float64 {
	return math.Pow(a+b, float64(n))
}
