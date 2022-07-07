package main_test

import (
	"reflect"
	"testing"

	"github.com/kshamiev/urok/sample/excel/typs"
)

// go test ./tutorial/gorutine/workerpool -run=^# -bench=WorkerPool -benchtime=1000000x -count 5 -cover -v
// 1000 1 5
// Benchmark_WorkerPool-8   	       4	 313749398 ns/op	  853638 B/op	   60513 allocs/op
// 1000 10 50
// Benchmark_WorkerPool-8   	     104	  10570623 ns/op	  801617 B/op	   59788 allocs/op
// 10000 100 500
// Benchmark_WorkerPool-8   	     120	   8580492 ns/op	  803321 B/op	   59791 allocs/op
// 10000 100 1000
// Benchmark_WorkerPool-8   	     127	   9083956 ns/op	  803158 B/op	   59791 allocs/op
// Benchmark_WorkerPool-8   	       4	 293293199 ns/op	  852010 B/op	   60476 allocs/op
func Benchmark_WorkerPool(b *testing.B) {
	obj := &typs.Cargo{
		ID:   "$%^",
		Name: "Popcorn",
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tt := reflect.TypeOf(obj)
		if tt.Kind() == reflect.Ptr {
			tt = tt.Elem()
		}
		_ = tt.String()
	}
}

// Benchmark_WorkerPool-8   	 7447424	       147.8 ns/op	     144 B/op	       2 allocs/op
// Benchmark_WorkerPool-8   	14 354 996	        73.38 ns/op	      80 B/op	       1 allocs/op
