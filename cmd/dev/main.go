package main

import (
	"fmt"
	_ "fmt"
)

type Rrrr struct {
}

func main() {

	test()

}

func test() {

	for i := 0; i < 100; i++ {
		var h *Rrrr
		if h == nil {
			fmt.Println("OK")
			h = &Rrrr{}
		}
	}

}
