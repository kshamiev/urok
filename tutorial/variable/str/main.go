package main

import (
	"fmt"
)

var Stroka0 string
var Stroka2 = "Строка pkg"

func main() {
	Stroka1 := "Строка local"
	Stroka0 = "qwerty"
	fmt.Println(Stroka0)
	fmt.Println(Stroka1)
	fmt.Println(Stroka2, "\n")

	sampleString(Stroka1)
	fmt.Println(Stroka1)

	sampleStringLink(&Stroka1)
	fmt.Println(Stroka1)
}

// передача по значению
func sampleString(s string) {

	s += " Add new string"

}

// передача по ссылке
func sampleStringLink(s *string) {

	*s += " Add new string"

}
