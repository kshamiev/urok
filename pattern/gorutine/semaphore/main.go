package main

import (
	"fmt"
	"time"
)

func main() {
	semaphor := make(chan bool, 5)
	for i := 0; i < 100; i++ {
		go func(in int) {
			semaphor <- true
			defer func() { <-semaphor }()
			fmt.Println(in)
			time.Sleep(time.Second)
		}(i)
	}
	var comIn string
	_, _ = fmt.Scanln(&comIn)
}
