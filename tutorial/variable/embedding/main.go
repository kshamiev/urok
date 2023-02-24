package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	g := Good{
		Model: Model{
			ID:    10,
			Name:  "Good",
			Price: decimal.NewFromFloat(100.50),
			Count: 10,
		},
		Capacity: 10,
		Width:    decimal.NewFromFloat(0.6),
		Length:   decimal.NewFromFloat(1.2),
		Height:   decimal.NewFromFloat(0.8),
		Weight:   decimal.NewFromFloat(50),
	}

	w := Ware{
		Model: Model{
			ID:    12,
			Name:  "Ware",
			Price: decimal.NewFromFloat(200.50),
			Count: 20,
		},
		Radius: decimal.NewFromFloat(80),
		Color:  "Red",
	}

	fmt.Println(g.Amount(), g.GetVolumeWeight())
	fmt.Println(w.Amount(), w.GetCircle())

}

type Model struct {
	ID    int64
	Name  string
	Price decimal.Decimal
	Count int
}

func (o *Model) Amount() decimal.Decimal {
	return o.Price.Mul(decimal.NewFromInt(int64(o.Count)))
}

type Good struct {
	Model
	Capacity int
	Width    decimal.Decimal
	Length   decimal.Decimal
	Height   decimal.Decimal
	Weight   decimal.Decimal
	// ...
}

func (o *Good) GetVolumeWeight() decimal.Decimal {
	return o.Width.Mul(o.Height).Mul(o.Length).Mul(o.Weight)
}

type Ware struct {
	Model
	Radius decimal.Decimal
	Color  string
	// ...
}

func (o *Ware) GetCircle() decimal.Decimal {
	return o.Radius.Mul(decimal.NewFromFloat(3.14)).Mul(decimal.NewFromInt(2))
}
