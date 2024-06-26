package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			// если не сделать получение с проверкой
			// то после закрытия канала селект будет бесконечно сваливаться в секцию с закрытым каналом читаю из него пустое значение
			case msg1 := <-c1:
				fmt.Println(msg1)
			// если не сделать получение с проверкой
			// то после закрытия канала селект будет бесконечно сваливаться в секцию с закрытым каналом читаю из него пустое значение
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
