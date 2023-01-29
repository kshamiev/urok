package main

import (
	"fmt"
)

// Составное ограничение типа (вынос ограничение в отдельный тип) для многократного использования

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

	data3 := []uint64{
		0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100,
	}
	fmt.Printf("AVG: %T = %v\n", Avg(data3), Avg(data3))
}

// Numeric expresses a type constraint satisfied by any numeric type.
type Numeric interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 |
		int64 | float32 | float64 // | complex128
}

func Avg[T Numeric](list []T) T {
	var s T
	for i := range list {
		s += list[i]
	}
	s = s / T(len(list))
	return s
}
