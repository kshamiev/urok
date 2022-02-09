package main

import "strconv"

func main() {
}

func RLE(str string) string {

	var res string
	var n int
	var step rune

	for _, s := range str {
		if step != s {
			if step != '' {

			}
			res += string(step) + strconv.Itoa(n)
			step = s
			n = 1
		} else {
			n++
		}
	}
	return res
}
