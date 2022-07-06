package main

import (
	"fmt"
	"reflect"

	"github.com/kshamiev/urok/sample/excel/typs"
)

func main() {

	// obj := Test1{Name: "popcorn"}
	obj := &typs.Cargo{
		ID:   "$%^",
		Name: "Popcorn",
	}

	tt := reflect.TypeOf(obj)
	if tt.Kind() == reflect.Ptr {
		tt = tt.Elem()
	}
	fmt.Println(tt.Kind())
	fmt.Println(tt.String())
	fmt.Println(tt.PkgPath())

}

var store = make(map[interface{}]bool)

type Test1 struct {
	Name string
}
