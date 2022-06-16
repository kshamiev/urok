package main

import (
	"fmt"
	. "fmt"
	_ "fmt"
)

func main() {

	fmt.Println(10000 / 500)

	f := Test()
	if f == nil {
		fmt.Println("NIL")
	} else {
		Println("NOT NIL")
	}

}

func Test() Face {
	// return nil
	// var g *Fikus
	// return g
	return Fikus{}
}

type Face interface {
	Good()
}

type Fikus struct {
}

func (o Fikus) Good() {

}
