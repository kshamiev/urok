// +build ignore go1.7 // запрет компиляции с указанием версии go
// package routing // import "application/routing" // указывает как импортировать пакет и где он должен находится
// точка входа в программу
package main

import (
	"fmt"
	_ "urok/pkg"
)

func init() {
	fmt.Println("main")
}

func main() {

}
