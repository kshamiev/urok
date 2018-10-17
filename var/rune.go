package main

import (
	"fmt"
)

func main() {

	Вася := 5
	Петя := 10

	Муся := Вася + Петя
	fmt.Println(Муся)
	fmt.Println()

	// Строка
	str := "Привет, Мир! 你好世界"
	fmt.Println("string: ", str, len(str))
	for i, v := range str {
		if v == 'в' {
			fmt.Println("начинается с байта ", i, " OK")
		} else if v == '世' {
			fmt.Println("начинается с байта ", i, " OK")
		}
		fmt.Printf("%v - %#U at position %d\n", v, v, i)
	}
	fmt.Println()

	// Байты
	bin := []byte(str)
	fmt.Println("binary: ", bin, len(bin))
	for i, v := range bin {
		fmt.Printf("raw binary index: %v, oct: %v, hex: %x\n", i, v, v)
	}
	fmt.Println()

	// Руны
	runes := []rune(str)
	fmt.Println("runes: ", runes, len(runes))
	for i, v := range runes {
		if v == 'в' {
			fmt.Println("начинается с позиции ", i, " OK")
		} else if v == '世' {
			fmt.Println("начинается с позиции ", i, " OK")
		}
		fmt.Printf("%v - %#U at position %d\n", v, v, i)
	}
	fmt.Println()

}
