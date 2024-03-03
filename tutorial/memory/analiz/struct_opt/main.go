package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	aaa [2]bool // смещение байтов: 0
	bbb int32   // смещение байтов: 4
	ccc [2]bool // смещение байтов: 8
}

type Bar struct {
	aaa [2]bool // смещение байтов: 0
	ccc [2]bool // смещение байтов: 2
	bbb int32   // смещение байтов: 4
}

func main() {
	ff := Foo{}
	bb := Bar{}
	fmt.Println(unsafe.Sizeof(ff))
	fmt.Println(unsafe.Sizeof(bb))
	fmt.Printf("offsets of fields: aaa: %+v; bbb: %+v; ccc: %+v\n", unsafe.Offsetof(ff.aaa), unsafe.Offsetof(ff.bbb), unsafe.Offsetof(ff.ccc))
	fmt.Printf("offsets of fields: aaa: %+v; ccc: %+v; bbb: %+v\n", unsafe.Offsetof(bb.aaa), unsafe.Offsetof(bb.ccc), unsafe.Offsetof(bb.bbb))
}

// с помощью go run запускается main.go
//
// смещения полей: aaa: 0; bbb: 4; ccc: 8
// смещения полей: aaa: 0; ccc: 2; bbb: 4
