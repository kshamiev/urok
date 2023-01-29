package main

import "fmt"

// T
// generic - (Общий) Это символ или мета-тип, представляющий один или несколько конкретных типов или интерфейсов

// int | int16 | int32 | int64 ...
// это ограничение, указывающее, какие конкретные типы можно использовать.

func main() {
	// Среднее арифметическое
	data1 := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	fmt.Printf("AVG: %T = %v\n", Avg(data1), Avg(data1))

	data2 := []float64{
		1.34, 2.65, 3.91, 4.23, 5.87, 6.48, 7.82, 8.59, 9.34, 10.72,
	}
	fmt.Printf("AVG: %T = %v\n", Avg(data2), Avg(data2))
}

func Avg[T int | float64](list []T) T {
	var s T
	for i := range list {
		s += list[i]
	}
	s = s / T(len(list))
	return s
}
