package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Calc struct {
	DepositStart decimal.Decimal // начальный депозит
	Deposit      decimal.Decimal // сумма нового депозита (реинвест)
	Percent      decimal.Decimal // начисляемый процент в день для депозита
	Period       int             // период для которого рассчитываем результат
	InputDay     decimal.Decimal // приход в день со всех депозитов
}

// начальная сумма всех работающих депозитов не включена в вывод результата
func main() {
	obj := &Calc{
		DepositStart: decimal.NewFromFloat(11000000),
		Deposit:      decimal.NewFromFloat(250000),
		Percent:      decimal.NewFromFloat(0.8),
		Period:       60,
	}
	obj.Calc()
}

func (cal *Calc) Calc() {
	cal.InputDay = cal.DepositStart.Div(decimal.NewFromInt(100)).Mul(cal.Percent)
	fmt.Printf("START: balance: 0 incominday: %s\n", cal.InputDay.String())

	var sum decimal.Decimal
	for i := 1; i <= cal.Period; i++ {
		sum = sum.Add(cal.InputDay)
		fmt.Printf("balance: %s\n", sum.Floor().String())

		// проверяем накопилась ли достаточная сумма для нового депозита
		if sum.GreaterThanOrEqual(cal.Deposit) {
			cal.DepositStart = cal.DepositStart.Add(cal.Deposit)
			sum = sum.Sub(cal.Deposit)
			// увеличение доходности в день за вложенный депозит
			cal.InputDay = cal.InputDay.Add(cal.Deposit.Div(decimal.NewFromInt(100)).Mul(cal.Percent))
			fmt.Printf("balance: %s incominday: %s day: %d\n", sum.Floor().String(), cal.InputDay.Floor().String(), i)
		}
	}
	fmt.Printf("balance full: %s \n", cal.DepositStart.String())
}
