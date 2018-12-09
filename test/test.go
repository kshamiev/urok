package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	f := "9000 / ( 720 + 4 ) * 3.5"
	result := Formula(f)
	fmt.Printf("formula: %s\nresult: %f", f, result)
}

func Formula(f string) (res float64) {
	f = strings.TrimSpace(f)
	res, _ = formula(strings.Split(f, " "))
	return
}

func formula(data []string) (res float64, level int) {
	for i := 0; i < len(data); i++ {
		if len(data) < i+2 {
			break
		}
		switch data[i] {
		case ")":
			return res, i
		case "/":
			i++
			if data[i] == "(" {
				i++
				r, n := formula(data[i:])
				res = res / r
				i = i + n
			} else if fl, err := strconv.ParseFloat(data[i], 64); err == nil {
				res = res / fl
			}
		case "*":
			i++
			if data[i] == "(" {
				i++
				r, n := formula(data[i:])
				res = res / r
				i = i + n
			} else if fl, err := strconv.ParseFloat(data[i], 64); err == nil {
				res = res * fl
			}
		case "+":
			i++
			if data[i] == "(" {
				i++
				r, n := formula(data[i:])
				res = res / r
				i = i + n
			} else if fl, err := strconv.ParseFloat(data[i], 64); err == nil {
				res = res + fl
			}
		case "-":
			i++
			if data[i] == "(" {
				i++
				r, n := formula(data[i:])
				res = res / r
				i = i + n
			} else if fl, err := strconv.ParseFloat(data[i], 64); err == nil {
				res = res - fl
			}
		default:
			if fl, err := strconv.ParseFloat(data[i], 64); err == nil {
				res = fl
			}
		}
	}
	return
}
