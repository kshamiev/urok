package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

// 0xcc 0x86
func main() {
	var f []byte

	f = []byte{0314, 0206}
	fmt.Println("%" + string(f) + "%")

	// src := []byte("48656c6c6f20476f7068657221")
	src := []byte("d398")
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%" + fmt.Sprintf("%s", dst[:n]) + "%")
}

//
// //  0xd2 0xb1
// f = []byte{0322, 0261}
// fmt.Println("%" + string(f) + "%")
//
// //  0xd2 0xb1
// f = []byte{0322, 0261}
// fmt.Println("%" + string(f) + "%")
//
// //  0xcb 0x9a
// f = []byte{0313, 0232}
// fmt.Println("%" + string(f) + "%")
//
// //  0xe2 0x80 0xa8
// f = []byte{0342, 0200, 0250}
// fmt.Println("%" + string(f) + "%")
//
// // 0xcc 0x86
// f = []byte{0314, 0206}
// fmt.Println("%" + string(f) + "%")
//
// // 0xd3 0x98
// f = []byte{0323, 0230}
// fmt.Println("%" + string(f) + "%")
//
// f = []byte{0313, 0232}
// fmt.Println("%" + string(f) + "%")
//
// // 0xe2 0x80 0xa8
// f = []byte{0342, 0200, 0250}
// fmt.Println("%" + string(f) + "%")
//
// // 0xef 0xbf 0xbd
// f = []byte{0357, 0277, 0275}
// fmt.Println("%" + string(f) + "%")
//
// // 0xcb 0x9c
// f = []byte{0313, 0234}
// fmt.Println("%" + string(f) + "%")
//
// // c3 9c
// f = []byte{0303, 0234}
// fmt.Println("%" + string(f) + "%")
//
// // 0xe2 0x82 0xbd
// f = []byte{0342, 0202, 0275}
// fmt.Println("%" + string(f) + "%")
