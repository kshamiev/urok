package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05Z"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	obj := &Item{
		ID: 23,
	}

	fmt.Println(obj)
	fmt.Println(obj.GetID())

}

type Item struct {
	Model
	ID int64
}

type Model struct {
	ID int64
}

func (m Model) GetID() int64 {
	return m.ID
}
