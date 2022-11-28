package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "codegen/doc.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}
	for _, example := range f.Comments {
		fmt.Println(example.Text())
	}
}
