package main

import (
	"fmt"
)

func main() {

	dd := make([]int, 3, 9)

	fmt.Println(len(dd), "OK")

	x := [6]string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(x[2:5])

}
