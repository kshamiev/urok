package main

import (
	"fmt"
	"time"
)

func main() {
	semaphor := make(chan bool, 5)
	for i := 0; i < 100; i++ {
		semaphor <- true
		go func(in int) {
			defer func() { <-semaphor }()
			fmt.Println(in)
			time.Sleep(time.Second * 5)
		}(i)
	}
}
