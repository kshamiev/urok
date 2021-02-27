// Порядки сложности алгоритма
// C – константа		// постоянная
// log(N)				// с логарифмическим понижающим коэффициентом
// N					// линейно от количества N
// N^C, C>1				// вложенные и зависимые циклы от N
// C^N, C>1				// рекурсия (C - количество рекурсий)
// N(N-1)! (факториал)	// N(N-1)! (3! = 6, 4! = 24, 6! = 720, ...)

// go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out bench_test.go ... > namefile.txt
// go test -bench=.

// Сложность алгоритма зависит от:
// 1) Объёма входных данных. // Как зависит скорость работы программы от количества входных данных
package testb

import (
	"testing"
)

// O(N^2)
func BenchmarkAlgorithmComplexity(b *testing.B) {
	data := GetAlgorithmComplexity1(1000)
	b.ResetTimer()
	b.ReportAllocs()
	res := AlgorithmComplexity1(1000, data)
	b.Log(len(res))
}
