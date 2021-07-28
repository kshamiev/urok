package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Item struct {
	ID      uint64
	Name    string
	Content string
	Price   float64
}

func main() {

	var result = make(map[uint64]Item)
	for i := uint64(0); i < 10; i++ {
		res := Item{
			ID:      45 + uint64(i),
			Name:    "Popcorn",
			Content: "Data Data Data",
			Price:   65.34 + float64(i),
		}
		result[i] = res
	}

	var data []byte
	var err error
	if data, err = json.Marshal(result); err != nil {
		panic(err)
	}

	logsSave(string(data))

	// var result = []Item{}
	// for i:=0; i < 10; i++ {
	//	res := Item{
	//		ID:      45 + uint64(i),
	//		Name:    "Popcorn",
	//		Content: "Data Data Data",
	//		Price:   65.34 + float64(i),
	//	}
	//	result = append(result, res)
	// }
	//
	//
	// var data []byte
	// var err error
	// if data, err = json.Marshal(result); err != nil {
	//	panic(err)
	// }
	//
	// logsSave(string(data))

}

var fp *os.File

// logsSave непосредственное сохранение лога
func logsSave(msg string) {

	_, pathFile, _, _ := runtime.Caller(0)

	pathFile = filepath.Dir(pathFile) + "/file.json"

	if fp == nil {
		_ = os.MkdirAll(filepath.Dir(pathFile), 0777)
		fp, _ = os.OpenFile(pathFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	}
	if fp != nil {
		_, _ = fp.WriteString(msg + "\n")
	}
	fmt.Println(msg)
}
