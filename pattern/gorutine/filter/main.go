package main

import (
	"fmt"
	"time"
)

func main() {
	chSet := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			chSet <- i
			time.Sleep(time.Second)
		}
		close(chSet)
	}()

	operator := func(n int) bool {
		return n%2 == 0
	}

	chGet := Filter(chSet, operator)

	for v := range chGet {
		fmt.Println(v)
	}
}

func Filter(inputStream <-chan int, operator func(int) bool) <-chan int {
	filteredStream := make(chan int)
	go func() {
		var i int
		var ok bool
		for {
			if i, ok = <-inputStream; !ok {
				close(filteredStream)
				return
			} else if !operator(i) {
				continue
			}
			filteredStream <- i
		}
	}()
	return filteredStream
}
