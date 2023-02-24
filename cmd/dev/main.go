package main

import "fmt"

func main() {

	fmt.Println("OK")

}

type Item struct {
	Model
	ID int64
}

type Model int
