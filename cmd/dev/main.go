package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	obj := Model(0)

	d, _ := json.Marshal(obj)

	fmt.Println(string(d))
	fmt.Println(d)

}

type Item struct {
	Model
	ID int64
}

type Model int
