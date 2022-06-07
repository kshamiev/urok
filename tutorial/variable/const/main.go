package main

import (
	"fmt"
)

const (
	_         = iota
	KB uint64 = 1 << (10 * iota) // тысяча
	MB                           // миллион
	GB                           // миллиард
	TB                           // триллион
	PB                           // квадриллион
	EB                           // квинтиллион
)

const (
	one = iota + 1
	two
	_
	foo
)

const number uint64 = 1 << 63
const SAMPLE float64 = 45.56
const SAMPLE1 = float64(45.56)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	fmt.Println()

	fmt.Printf("%b\n", number)
	fmt.Printf("%o\n", number)
	fmt.Printf("%d\n", number)
	fmt.Printf("%x\n", number)
	fmt.Println()

	fmt.Println(one, two, foo)
	fmt.Println(SAMPLE, SAMPLE1)
}
