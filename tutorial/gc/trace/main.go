package main

import (
	"time"
)

var ballast []byte

// go build
// GODEBUG=gctrace=1 ./trace
func main() {
	// ballast = make([]byte, 10<<30)
	add(100000)
	add(100000)
	add(100000)
	add(100000)
	add(100000)
	time.Sleep(time.Hour)
}

var numbers []int

func add(n int) {
	for i := 0; i < n; i++ {
		numbers = append(numbers, i+1)
	}
}
