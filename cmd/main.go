package main

import (
	"fmt"
)

func main() {
	var f []byte

	f = []byte{0313, 0232}
	fmt.Println("%" + string(f) + "%")

	// 0xe2 0x80 0xa8
	f = []byte{0342, 0200, 0250}
	fmt.Println("%" + string(f) + "%")

	// 0xef 0xbf 0xbd
	f = []byte{0357, 0277, 0275}
	fmt.Println("%" + string(f) + "%")

	// 0xcb 0x9c
	f = []byte{0313, 0234}
	fmt.Println("%" + string(f) + "%")

	// c3 9c
	f = []byte{0303, 0234}
	fmt.Println("%" + string(f) + "%")

	// 0xe2 0x82 0xbd
	f = []byte{0342, 0202, 0275}
	fmt.Println("%" + string(f) + "%")
}
