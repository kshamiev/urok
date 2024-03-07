package main

import (
	"fmt"
	"io"
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
	f, e := os.Open("cmd/dev/example.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	decoder := charmap.Windows1251.NewDecoder()
	reader := decoder.Reader(f)
	b, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
