package function

import "fmt"

type FuncParameter func(int, int)

func Parameter(n, m int) {
	fmt.Println(n + m)
}

func ParameterSum1(x, y int, n FuncParameter) {
	n(x, y)
}

func ParameterSum2(x, y int, n func(int, int) int) (int, bool) {
	res := n(x, y)
	if res == 35 {
		return res, true
	}
	return res, false
}
