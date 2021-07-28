// Порядки сложности алгоритма - оценка порядка сложности
// C – константа		// постоянная
// log(N)				// с логарифмическим понижающим коэффициентом
// N					// линейно от количества N
// N^C, C>1				// вложенные и зависимые циклы от N (C - количество вложенных циклов)
// C^N, C>1				// рекурсия (C - количество рекурсий)
// N(N-1)! (факториал)	// N(N-1)! (3! = 6, 4! = 24, 6! = 720, ...)

// Сложность алгоритма зависит от:
// 1) Объёма входных данных. // Как зависит скорость работы программы от количества входных данных
package testb

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

// O(N^2)
func BenchmarkAlgorithmComplexity(b *testing.B) {
	data := GetAlgorithmComplexity1(1000)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		res := AlgorithmComplexity1(1000, data)
		b.Log(len(res))
	}
}

// ////

func BenchmarkSample(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
	b.Log("OK")
}

// ////

type Set struct {
	set map[interface{}]struct{}
	mu  sync.Mutex
}

func (s *Set) Add(x interface{}) {
	s.mu.Lock()
	s.set[x] = struct{}{}
	s.mu.Unlock()
}

func (s *Set) Delete(x interface{}) {
	s.mu.Lock()
	delete(s.set, x)
	s.mu.Unlock()
}

func BenchmarkSetDelete(b *testing.B) {
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		set := Set{set: make(map[interface{}]struct{})}
		for _, elem := range testSet {
			set.Add(elem)
		}
		b.StartTimer()
		for _, elem := range testSet {
			set.Delete(elem)
		}
	}
}
