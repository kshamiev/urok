package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	obj := &Calc{
		Percent: decimal.NewFromFloat(0.8),
		Deposit: decimal.NewFromFloat(200000),

		InputDay: decimal.NewFromFloat(30000),
	}

	_ = obj.Calc()
}

type Calc struct {
	Percent decimal.Decimal // начисляемый процент в день для депозита
	Deposit decimal.Decimal // минимальная сумма депозита

	InputDay decimal.Decimal // приход в день
}

func (cal *Calc) Calc() decimal.Decimal {
	var sum decimal.Decimal

	fmt.Printf("START: balance: 0 incominday: %s\n", cal.InputDay.String())

	for i := 1; i <= 100; i++ {
		sum = sum.Add(cal.InputDay)
		fmt.Printf("balance: %s\n", sum.Floor().String())

		// проверяем накопилась ли достаточная сумма для нового депозита
		if sum.GreaterThanOrEqual(cal.Deposit) {
			// увеличение доходности в день за вложенный депозит
			cal.InputDay = cal.InputDay.Add(sum.Div(decimal.NewFromInt(100)).Mul(cal.Percent))
			// вкладываем все
			sum = decimal.Decimal{}
			fmt.Printf("balance: %s incominday: %s day: %d\n", sum.Floor().String(), cal.InputDay.Floor().String(), i)
		}
	}

	return sum
}

// 250 000 0.8
// balance: 40000 incominday: 64000 day: 98
// balance: 0 incominday: 64935 day: 99

// 30 000 0.7
// balance: 1039500 incominday: 51000
// balance: 0 incominday: 60265 day: 100
