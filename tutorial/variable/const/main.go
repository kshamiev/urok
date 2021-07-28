package main

import (
	"fmt"
)

const (
	_         = iota
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

const (
	one = iota + 1
	two
	_
	foo
)

const SAMPLE float64 = 45.56
const SAMPLE1 = float64(45.56)

func main() {

	fmt.Println(KB, MB, GB, TB, PB, EB, "\n")
	fmt.Println(one, two, foo, "\n")
	fmt.Println(SAMPLE, SAMPLE1)

}
