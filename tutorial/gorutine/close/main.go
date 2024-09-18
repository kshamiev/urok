package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 20)
	ch <- "echo"
	ch <- "case,"
	ch <- "if,"
	close(ch)

	for i := 0; i < 10; i++ {
		// for v := range ch {
		// 	fmt.Println("---" + v + "---")
		// }

		fmt.Println(<-ch)

		select {
		// case val := <-ch:
		// 	fmt.Println("select", val)
		case val, ok := <-ch:
			fmt.Println("select", val, ok)
		}

		if val, ok := <-ch; ok {
			fmt.Println(val, ok)
		} else {
			break
		}
	}
	fmt.Println("finish")
}
