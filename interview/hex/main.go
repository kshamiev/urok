package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Нужно реализовать сложение больших 16 разрядных чисел
func main() {
	s, err := sum("44a2e2a50e24459bb771e2e4b801c55087894b9aa72e", "bc2d840be66ff0ee7a1a")
	fmt.Println(s, err)
	// 44a2e2a50e24459bb771e2e5742f495c6df93c892148 nil
	s, err = sum("bc2d840be66ff0ee7a1a", "44a2e2a50e24459bb771e2e4b801c55087894b9aa72e")
	fmt.Println(s, err)
	// 44a2e2a50e24459bb771e2e5742f495c6df93c892148 nil
}

func sum(x, y string) (string, error) {
	xLen := len(x)
	yLen := len(y)
	if yLen > xLen {
		xLen, yLen = yLen, xLen
		x, y = y, x
	}
	var x1, y1 int64
	var err error
	var flagOver bool
	resList := make([]string, xLen)
	xLen--
	yLen--

	for i := xLen; i > -1; i-- {
		x1, err = strconv.ParseInt(string(x[i]), 16, 64)
		if err != nil {
			return "", err
		}
		if yLen >= 0 {
			y1, err = strconv.ParseInt(string(y[yLen]), 16, 64)
			if err != nil {
				return "", err
			}
			x1 += y1
			yLen--
		}
		if flagOver {
			x1++
			flagOver = false
		}
		if x1 > 15 {
			x1 = x1 - 16
			flagOver = true
		}
		resList[i] = strconv.FormatInt(x1, 16)
	}
	if flagOver {
		return "1" + strings.Join(resList, ""), nil
	} else {
		return strings.Join(resList, ""), nil
	}
}
