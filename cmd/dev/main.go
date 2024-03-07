package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	// Запись строки в кодировке Windows-1252
	encoder := charmap.Windows1251.NewEncoder()
	s, e := encoder.String("Распоряжения")
	if e != nil {
		log.Fatal(e)
	}
	_ = os.WriteFile("cmd/dev/example.txt", []byte(s), os.ModePerm)

	// Декодировка в UTF-8
	f, e := os.ReadFile("cmd/dev/example.txt")
	if e != nil {
		log.Fatal(e)
	}
	decoder := charmap.Windows1251.NewDecoder()
	b, e := decoder.Bytes(f)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(b))
}
