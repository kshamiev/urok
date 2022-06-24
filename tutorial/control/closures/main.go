package main

import "fmt"

func main() {
	name := "popcorn"

	closures := func() {
		name = "замыкание рулит"
	}

	fmt.Println(name)
	closures()
	fmt.Println(name)
}
