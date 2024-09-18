package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	for {
		_, _ = fmt.Scanln(&input)
		if input == "exit" {
			break
		}
	}
}
