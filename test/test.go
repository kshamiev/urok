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

const (
	formula_skopen  = "("
	formula_skclose = ")"
	formula_minus   = "-"
	formula_plus    = "+"
	formula_um      = "*"
	formula_del     = "/"
)

func Formula(f string) (res float64) {
	f = strings.TrimSpace(f)
	res, _ = formula(strings.Split(f, " "))
	return
}

func formula(data []string) (res float64, level int) {
	var (
		err error
		r   float64
		n   int
	)
	for i := 0; i < len(data); i++ {
		if len(data) < i+2 {
			break
		}
		switch data[i] {
		case formula_skclose:
			return res, i
		case formula_del:
			i++
			if data[i] == formula_skopen {
				i++
				r, n = formula(data[i:])
				res = res / r
				i = i + n
			} else if r, err = strconv.ParseFloat(data[i], 64); err == nil {
				res = res / r
			}
		case formula_um:
			i++
			if data[i] == formula_skopen {
				i++
				r, n = formula(data[i:])
				res = res / r
				i = i + n
			} else if r, err = strconv.ParseFloat(data[i], 64); err == nil {
				res = res * r
			}
		case formula_plus:
			i++
			if data[i] == formula_skopen {
				i++
				r, n = formula(data[i:])
				res = res / r
				i = i + n
			} else if r, err := strconv.ParseFloat(data[i], 64); err == nil {
				res = res + r
			}
		case formula_minus:
			i++
			if data[i] == formula_skopen {
				i++
				r, n = formula(data[i:])
				res = res / r
				i = i + n
			} else if r, err := strconv.ParseFloat(data[i], 64); err == nil {
				res = res - r
			}
		default:
			if fl, err := strconv.ParseFloat(data[i], 64); err == nil {
				res = fl
			}
		}
	}
	return
}
