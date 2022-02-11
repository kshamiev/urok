package main

import "fmt"

func main() {
	f := []byte{
		0303,
		0234,
	} // "Ü"
	f = []byte{
		0313,
		0232,
	} // "˚"
	f = []byte{ // 0xe2 0x80 0xa8
		0342,
		0200,
		0250,
	} // "% %"
	fmt.Println("%" + string(f) + "%")
}
