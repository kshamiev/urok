package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for n := 0; n < 10; n++ {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
		close(c1)
	}()

	go func() {
		for n := 0; n < 10; n++ {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
		close(c2)
	}()

	go func() {
		for {
			select {
			// если не сделать получение с проверкой
			// то после закрытия канала селект будет бесконечно сваливаться в секцию с закрытым каналом читаю из него пустое значение
			case msg1, ok := <-c1:
				if !ok {
					goto Poit
				}
				fmt.Println(msg1)
			// если не сделать получение с проверкой
			// то после закрытия канала селект будет бесконечно сваливаться в секцию с закрытым каналом читаю из него пустое значение
			case msg2, ok := <-c2:
				if !ok {
					goto Poit
				}
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	Poit:
	}()

	var input string
	fmt.Scanln(&input)
}
