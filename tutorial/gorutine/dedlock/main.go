// deadlock это когда мы бесконечно ждем чтения из канала и не можем прочитать
// (в канал никто не пишет)
// или наоборот бесконечно пишем в канал который никто не будет читать.
// Это справедливо в работе основной горутине (main).
// В порожденных (дочерних) горутинах дедлока не будет. Так порождаются зомбо процессы. )))
package main

import (
	"fmt"
	"strconv"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping: " + strconv.Itoa(i)
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
	var c chan string = make(chan string)

	// c <- "ping"
	// fatal error: all goroutines are asleep - deadlock!
	// <-c
	// fatal error: all goroutines are asleep - deadlock!

	go pinger(c)
	go printer(c)

	var input string
	_, _ = fmt.Scanln(&input)
	fmt.Println(input)
}
