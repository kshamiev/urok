// sample full
// go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out bench_test.go ... > namefile.txt
//
// go test -bench=. > old.txt
// change
// go test -bench=. > new.txt
// benchstat -html old.txt new.txt > diff.html
package test

import (
	"fmt"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}

func BenchmarkToJSON(b *testing.B) {
	tmp := &testStruct{X: 1, Y: "string"}
	js, err := tmp.ToJSON()
	if err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.SetBytes(int64(len(js)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := tmp.ToJSON(); err != nil {
			b.Fatal(err)
		}
	}
}
