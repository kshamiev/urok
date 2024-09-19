package main

/* producer-consumer problem in Go */

import (
	"fmt"
)

func produceGo() <-chan int {
	msgCH := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			msgCH <- i
		}
		close(msgCH)
	}()
	return msgCH
}

func consumeGo(ch <-chan int) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		for msg := range ch {
			fmt.Println(msg)
		}
		done <- struct{}{}
	}()
	return done
}

func main() {
	msgCH := produceGo()
	done := consumeGo(msgCH)
	<-done
}
