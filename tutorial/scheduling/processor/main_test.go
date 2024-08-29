package main

import (
	"runtime"
	"testing"
)

// GOGC=off go test ./tutorial/scheduling/sample/processor/. -run none -bench . -benchtime=10x -count 1 -cpu 1
// GOGC=off go test ./tutorial/scheduling/sample/processor/. -run none -bench . -benchtime=10x -count 1 -cpu 16

var numbers = generateList(1e9)

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(numbers)
	}
}

func BenchmarkConcurrentOrParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addConcurrentOrParallel(runtime.NumCPU(), numbers)
	}
}

// Пример пузырьковой сортировки или подобных алгоритмов не подходит для разделения вычислений на несколько процессов.
// В силу самого алгоритма это не имеет смысла. Так как в конце концов будет произведена повторно полная сортировка всех частей исходных данных
