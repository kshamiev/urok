package workerpool

import (
	"fmt"
	"testing"
	"time"
)

type ExampleTask string

func (e ExampleTask) Execute() {
	time.Sleep(time.Microsecond * 10)
}

func TestNewPool(t *testing.T) {
	pool := NewPool(1000, 1, 20)
	pool.TaskAdd(ExampleTask("foo"))
	pool.TaskAdd(ExampleTask("bar"))
	for i := 0; i < 2000000; i++ {
		pool.TaskAdd(ExampleTask(fmt.Sprintf("additional_%d", i+1)))
	}
	pool.Wait()
}

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
	pool := NewPool(1000, 1, 5)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pool.TaskAdd(ExampleTask("foo"))
		pool.TaskAdd(ExampleTask("bar"))
		for i := 0; i < 20000; i++ {
			pool.TaskAdd(ExampleTask(fmt.Sprintf("additional_%d", i+1)))
		}
	}
	pool.Wait()
}
