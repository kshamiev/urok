package raszdel1

import (
	"math"
	"testing"
)

func TestPythagoras(t *testing.T) {
	c := Pythagoras(2, 3)
	t.Log(c)
	t.Log(math.Sqrt(25))
}

func TestEuclid(t *testing.T) {
	c := Euclid(8, 5)
	t.Log(c)

	c = Euclid(53, 21)
	t.Log(c)

	c = Euclid(math.Sqrt(2), 1)
	t.Log(c)

	c = Euclid(2.5, 1)
	t.Log(c)

	c = Euclid(24, 6)
	t.Log(c)

	c = Euclid(24, 18)
	t.Log(c)
}

func TestGaus(t *testing.T) {
	c := Gaus(100)
	t.Log(c)
}

func TestSquareOfTheNumber(t *testing.T) {
	c := SquareOfTheNumber(1, 3, 5)
	t.Log(c)
	c = SquareOfTheNumber(1, 3, 5, 7)
	t.Log(c)
}

func TestBenom(t *testing.T) {
	c := Benom(2, 3, 3)
	t.Log(c)
}
