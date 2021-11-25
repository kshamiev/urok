// sample full
// go test -bench=. bench_test.goq ... > namefile.txt
// go test -bench=BenchmarkToJSON bench_test.go ... > namefile.txt
//
// go test -bench=. > old.txt
// change method
// go test -bench=. > new.txt
// delta
// benchstat -html -sort name old.txt new.txt > delta.html
package delta

import (
	"testing"
)

// old BenchmarkToJSON-8   	 5907133	       210 ns/op	  95.28 MB/s	      32 B/op	       1 allocs/op
// new BenchmarkToJSON-8   	31448481	       38.6 ns/op	 596.54 MB/s	       0 B/op	       0 allocs/op
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
