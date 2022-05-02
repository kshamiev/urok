package main

import "fmt"

type Face interface {
	Good()
}

type Fikus struct {
}

func (o Fikus) Good() {

}

func main() {

	f := Test()
	if f == nil {
		fmt.Println("NIL")
	} else {
		fmt.Println("NOT NIL")
	}

}

func Test() Face {

	return nil
	// var g *Fikus
	// return g
	// return &Fikus{}
}
