package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"math/big"
	"reflect"

	"github.com/kshamiev/urok/sample/excel/typs"
)

func GenInt(x int64) int64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(x))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}

func main() {

	fmt.Println(GenInt(10000))

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

	store["gfdgfd"] = []interface{}{
		obj,
	}

	Dumper(store)

}

var store = make(map[string][]interface{})

type Test1 struct {
	Name string
}

// Dumper all variables to STDOUT
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
