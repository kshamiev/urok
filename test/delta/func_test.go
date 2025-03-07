// sample full
// go test -bench=. -benchmem
// go test -bench=. bench_test.goq ... > namefile.txt
// go test -bench=BenchmarkToJSON bench_test.go ... > namefile.txt
//
// go test -bench=. -benchmem -benchtime=1000000x -count 5 > old.txt
// go test -bench=. -benchmem -benchtime=10s -count 5 > old.txt
// change method
// go test -bench=. -benchmem -benchtime=1000000x -count 5 > new.txt
// go test -bench=. -benchmem -benchtime=10s -count 5 > new.txt
// delta
// benchstat -html -sort name old.txt new.txt > delta.html
// benchcmp old.txt new.txt
//
// go build -gcflags=-m
// показывает куда будут аллоцированы переменные программы
package delta

import (
	"testing"
)

func BenchmarkToJSON(b *testing.B) {
	tmp := &testStruct{X: 1, Y: "string"}
	for i := 0; i < b.N; i++ {
		if _, err := tmp.ToJSON(); err != nil {
			b.Fatal(err)
		}
	}
}
