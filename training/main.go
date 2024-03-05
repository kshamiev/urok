package main

import (
	"fmt"
	"math"

	"github.com/kshamiev/urok/training/stepik/raszdel1"
)

func main() {
	fmt.Println(raszdel1.Euclid(53, 21))
	fmt.Println(raszdel1.Euclid(math.Sqrt(2), 1))
	fmt.Println(raszdel1.Gaus(100))
	fmt.Println(raszdel1.SquareOfTheNumber(11))
	fmt.Println(raszdel1.Benom(2, 3, 3))
}
