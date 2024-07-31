package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var percent = decimal.NewFromFloat(0.20)
var add = decimal.NewFromFloat(250000)
var year = 3

func main() {
	fmt.Println("выплаты раз в месяц:")
	month()
	fmt.Println("")

	fmt.Println("выплаты раз в квартал:")
	quarter()
	fmt.Println("")

	fmt.Println("выплаты раз в полгода:")
	polgoda()
	fmt.Println("")
}

func month() {
	period := decimal.NewFromFloat(12)
	p := decimal.Decimal{}
	cash := decimal.Decimal{}
	for i := 0; i < year; i++ {
		for j := 0; j < 12; j++ {
			cash = cash.Add(add)
			if !p.IsZero() {
				cash = cash.Add(p)
			}
			p = cash.Mul(percent).Div(period)
		}
	}
	cash = cash.Add(p)
	p = cash.Mul(percent).Div(period)
	fmt.Println("капитал: ", cash.IntPart())
	fmt.Println("доход в месяц: ", cash.Mul(percent).Div(decimal.NewFromFloat(12)).IntPart())
}

func quarter() {
	period := decimal.NewFromFloat(4)
	p := decimal.Decimal{}
	cash := decimal.Decimal{}
	for i := 0; i < year; i++ {
		for j := 0; j < 12; j++ {
			cash = cash.Add(add)
			if j%3 == 0 {
				if !p.IsZero() {
					cash = cash.Add(p)
				}
				p = cash.Mul(percent).Div(period)
			}
		}
	}
	cash = cash.Add(p)
	p = cash.Mul(percent).Div(period)
	fmt.Println("капитал: ", cash.IntPart())
	fmt.Println("доход в месяц: ", cash.Mul(percent).Div(decimal.NewFromFloat(12)).IntPart())
}

func polgoda() {
	period := decimal.NewFromFloat(2)
	p := decimal.Decimal{}
	cash := decimal.Decimal{}
	for i := 0; i < year; i++ {
		for j := 0; j < 12; j++ {
			cash = cash.Add(add)
			if j%6 == 0 {
				if !p.IsZero() {
					cash = cash.Add(p)
				}
				p = cash.Mul(percent).Div(period)
			}
		}
	}
	cash = cash.Add(p)
	p = cash.Mul(percent).Div(period)
	fmt.Println("капитал: ", cash.IntPart())
	fmt.Println("доход в месяц: ", cash.Mul(percent).Div(decimal.NewFromFloat(12)).IntPart())
}
