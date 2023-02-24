package function

import "fmt"

type FuncParameter func(int, int)

func Parameter(n, m int) {
	fmt.Println(n + m)
}

func ParameterSum1(x, y int, n FuncParameter) {
	n(x, y)
}

func ParameterSum2(x, y int, n func(int, int)) {
	n(x, y)
}
