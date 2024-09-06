package main

import (
	"fmt"
	"time"
)

// go build
// GODEBUG=gctrace=1 ./trace
func main() {
	fmt.Println("start")

	fmt.Println("1")
	add(100000)
	time.Sleep(time.Second)

	fmt.Println("2")
	add(100000)
	time.Sleep(time.Second)

	fmt.Println("3")
	add(100000)
	time.Sleep(time.Second)

	fmt.Println("4")
	add(100000)
	time.Sleep(time.Second)

	fmt.Println("5")
	add(100000)
	time.Sleep(time.Hour)
}

var numbers []int

func add(n int) {
	for i := 0; i < n; i++ {
		numbers = append(numbers, i+1)
	}
}
