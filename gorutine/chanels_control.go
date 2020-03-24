package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func main() {

	ch := make(chan string)
	close(ch)
	fmt.Println("ok")

	for i := 0; i < 10; i++ {
		data := <-ch
		Dumper(data)

		if _, ok := <-ch; !ok {
			break
		}

		// for data := range <-ch {
		// 	Dumper(data)
		// }

	}

	fmt.Println("ok")
}

// Dump all variables to STDOUT
// From local debug
func Dumper(idl ...interface{}) string {
	ret := dump(idl...)
	fmt.Print(ret.String())

	return ret.String()
}

// dump all variables to bytes.Buffer
func dump(idl ...interface{}) bytes.Buffer {
	var buf bytes.Buffer

	var wr = io.MultiWriter(&buf)

	for _, field := range idl {
		fset := token.NewFileSet()
		_ = ast.Fprint(wr, fset, field, ast.NotNilFilter)
	}

	return buf
}
