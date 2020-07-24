package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 5)
	go func() {
		time.Sleep(time.Second * 5)
		for z := range ch {
			fmt.Println(z)
		}
		fmt.Println("===")
	}()

	ch <- 4
	ch <- 5
	ch <- 6
	close(ch)
	fmt.Println("OK")
	time.Sleep(time.Second * 10)
	fmt.Println("OK")
}
