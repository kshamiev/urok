package main

import (
	"fmt"
)

func main() {
	// Строка
	str := "Привет, Мир! 你好世界"
	fmt.Printf("string: %s length: %d\n", str, len(str))
	for i, v := range str {
		if v == 'в' {
			fmt.Printf("буква 'в' начинается с байта %d\n", i)
		} else if v == '世' {
			fmt.Printf("буква '世' начинается с байта %d\n", i)
		}
		fmt.Printf("rune: %v at position %d\n", v, i)
	}
	fmt.Println()

	// Байты
	bin := []byte(str)
	fmt.Println("binary: ", bin, "length: ", len(bin))
	for i, v := range bin {
		fmt.Printf("raw binary index: %v, oct: %v, hex: %x\n", i, v, v)
	}
	fmt.Println()

	// Руны
	runes := []rune(str)
	fmt.Println("runes: ", runes, len(runes))
	for i, v := range runes {
		if v == 'в' {
			fmt.Printf("%v at position %d\n", v, i)
		} else if v == '世' {
			fmt.Printf("%v at position %d\n", v, i)
		}
		fmt.Printf("rune: %v at position %d\n", v, i)
	}
	fmt.Println()

	for index, runeValue := range str {
		fmt.Printf("Позиция '%d', руна: ---%#U--- [%d]\n", index, runeValue, runeValue)
	}
}
