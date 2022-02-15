package main

import (
	"fmt"
)

func main() {
	var f []byte
	f = []byte{
		0313,
		0232,
	} // "˚"
	f = []byte{ // 0xe2 0x80 0xa8
		0342,
		0200,
		0250,
	} // "% %"
	f = []byte{ // 0xef 0xbf 0xbd
		0357,
		0277,
		0275,
	} // %�%
	f = []byte{ // 0xcb 0x9c
		0313,
		0234,
	} // %˜%
	f = []byte{ // c3 9c
		0303,
		0234,
	} // "%Ü%"
	f = []byte{ // 0xe2 0x82 0xbd
		0342,
		0202,
		0275,
	} // %₽%
	fmt.Println("%" + string(f) + "%")
}
