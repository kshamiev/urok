package main

import (
	"runtime"
	"testing"
)

// GOGC=off go test ./tutorial/scheduling/sample/processor/. -run none -bench . -benchtime=10x -count 1 -cpu 1
// GOGC=off go test ./tutorial/scheduling/sample/processor/. -run none -bench . -benchtime=10x -count 1 -cpu 16

var docs = generateList(1e3)

func BenchmarkSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		find("test", docs)
	}
}

func BenchmarkConcurrentOrParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findConcurrentOrParallel(runtime.NumCPU(), "test", docs)
	}
}
