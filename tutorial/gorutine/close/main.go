package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 20)
	ch <- "popcorn"
	ch <- "qwerty"
	ch <- "vasya"
	close(ch)

	for i := 0; i < 10; i++ {
		// for v := range ch {
		// 	fmt.Println("---" + v + "---")
		// }

		fmt.Println("!!!" + <-ch + "!!!")

		select {
		case val, ok := <-ch:
			fmt.Println("+++"+val+"+++", ok)
		}

		if val, ok := <-ch; ok {
			fmt.Println("..." + val + "...")
		} else {
			break
		}
	}
	fmt.Println("ok")
}
