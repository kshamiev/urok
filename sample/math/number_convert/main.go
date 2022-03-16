package main

import (
	"fmt"
	"strconv"
)

func bin(i int, prefix bool) string {
	i64 := int64(i)

	if prefix {
		return "0b" + strconv.FormatInt(i64, 2) // base 2 for binary
	} else {
		return strconv.FormatInt(i64, 2) // base 2 for binary
	}
}

func bin2int(binStr string) int {

	// base 2 for binary
	result, _ := strconv.ParseInt(binStr, 2, 64)
	return int(result)
}

func oct(i int, prefix bool) string {
	i64 := int64(i)

	if prefix {
		return "0o" + strconv.FormatInt(i64, 8) // base 8 for octal
	} else {
		return strconv.FormatInt(i64, 8) // base 8 for octal
	}
}

func oct2int(octStr string) int {
	// base 8 for octal
	result, _ := strconv.ParseInt(octStr, 8, 64)
	return int(result)
}

func hex(i int, prefix bool) string {
	i64 := int64(i)

	if prefix {
		return "0x" + strconv.FormatInt(i64, 16) // base 16 for hexadecimal
	} else {
		return strconv.FormatInt(i64, 16) // base 16 for hexadecimal
	}
}

func hex2int(hexStr string) int {
	// base 16 for hexadecimal
	result, _ := strconv.ParseInt(hexStr, 16, 64)
	return int(result)
}

func main() {

	num1 := int64(123456789)
	fmt.Println(fmt.Sprintf("%d", num1))
	fmt.Println(fmt.Sprintf("%b", num1))
	fmt.Println(fmt.Sprintf("%o", num1))
	fmt.Println(fmt.Sprintf("%x", num1))

	num := 123456789
	fmt.Println("Integer : ", num)
	fmt.Println("Binary : ", bin(num, false))
	fmt.Println("Octal : ", oct(num, true))
	fmt.Println("Hex : ", hex(num, true))

	// bin2int function does not handle the prefix
	// so set second parameter to false
	// otherwise you will get funny result

	fmt.Println("Binary to Integer : ", bin2int(bin(num, false)))
	fmt.Println("Octal to Integer : ", oct2int(oct(num, false)))
	fmt.Println("Hexadecimal to Integer : ", hex2int(hex(num, false)))

}
