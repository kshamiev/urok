package main

import "fmt"

func main() {

	// 5 101
	// 2 10
	// 10 1010
	fmt.Println(5 & 2)  // 0
	fmt.Println(5 | 2)  // 7
	fmt.Println(5 ^ 2)  // 7
	fmt.Println(2 ^ 2)  // 0
	fmt.Println(3 ^ 3)  // 0
	fmt.Println(5 ^ 10) // 15

}
