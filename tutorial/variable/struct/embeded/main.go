package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func main() {
	obj := &MyControl{}
	obj.ID = 67
	obj.Name = "Popcorn"
	obj.Price = 56.78
	obj.Control.ID = 6767
	obj.Control.Name = "Popcorn control"
	obj.Control.Price = 5556.78

	obj.Calc()
	obj.Design()
	obj.Names()

	Dumper("")

	obj.Control.Calc()
	obj.Control.Design()
	obj.Control.Names()

}

type MyControl struct {
	ID    uint64
	Name  string
	Price float64
	Control
}

func (ctr *MyControl) Names() {
	fmt.Println("Names", ctr.ID, ctr.Name, ctr.Price)
	fmt.Println("Names", ctr.Control.ID, ctr.Control.Name, ctr.Control.Price)
}

type Control struct {
	ID    uint64
	Name  string
	Price float64
}

func (ctr *Control) Design() {
	fmt.Println("Design", ctr.ID, ctr.Name, ctr.Price)
}

func (ctr *Control) Calc() {
	fmt.Println("Calc", ctr.ID, ctr.Name, ctr.Price)
}

func (ctr *Control) Names() {
	fmt.Println("Names", ctr.ID, ctr.Name, ctr.Price)
}

// //////////////////////

// Dump all variables to STDOUT
func Dumper(idl ...interface{}) string {
	ret := dump(idl...)
	fmt.Print(ret.String())
	return ret.String()
}

// dump all variables to bytes.Buffer
func dump(idl ...interface{}) bytes.Buffer {
	var buf bytes.Buffer
	var wr io.Writer

	wr = io.MultiWriter(&buf)
	for _, field := range idl {
		fset := token.NewFileSet()
		_ = ast.Fprint(wr, fset, field, ast.NotNilFilter)
	}
	return buf
}
