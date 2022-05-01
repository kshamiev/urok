package testb

import (
	"fmt"
	"testing"
)

// go test -bench=. -benchmem -benchtime=1000000x -count 5
// go test -bench=. -benchmem -benchtime=10s -count 5
func Benchmark_StringConcat(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		leftpad1("popcorn", 30, 'q')
	}
}

func Benchmark_BytesBuffer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		leftpad2("popcorn", 30, 'q')
	}
}

func Benchmark_StringsBuilder(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		leftpad21("popcorn", 30, 'q')
	}
}

func Benchmark_StringsRepeat(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		leftpad3("popcorn", 30, 'q')
	}
}

// ////

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 74382},
	{input: 382399},
}

func BenchmarkPrimeNumbers(b *testing.B) {
	b.Skip()
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {

				// testFunc(v.input)
			}
		})
	}
}
