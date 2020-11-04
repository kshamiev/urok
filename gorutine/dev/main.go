package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	close(ch)
	fmt.Println("ok")
	for i := 0; i < 10; i++ {
		for d := range <-ch {
			fmt.Println("!!!", d, "!!!")
		}
		data := <-ch
		fmt.Println("!!! " + data + " !!!")
		if _, ok := <-ch; !ok {
			break
		}
	}
	fmt.Println("ok")
}
