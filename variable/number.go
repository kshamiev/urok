package main

import (
	"fmt"
)

var Celoe0 int64
var Celoe2 = int64(23)

var Drobnoe0 float64
var Drobnoe2 = float64(34.45)

func main() {
	Celoe1 := int64(23)
	Drobnoe1 := float64(34.45)

	sampleNumber(Celoe2, Drobnoe2)

	fmt.Println(Celoe0, Celoe1, Celoe2)
	fmt.Println(Drobnoe0, Drobnoe1, Drobnoe2, "\n")

	sampleNumberLink(&Celoe2, &Drobnoe2)

	fmt.Println(Celoe0, Celoe1, Celoe2)
	fmt.Println(Drobnoe0, Drobnoe1, Drobnoe2)
}

// передается по значению
func sampleNumber(x int64, y float64) {
	x += 24
	y += 4.5
}

// передается по ссылке
func sampleNumberLink(x *int64, y *float64) {
	*x += 24
	*y += 4.5
}
