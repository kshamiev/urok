package workerbus

import (
	"fmt"
	"testing"
)

type ExampleTask struct {
	Name string
}

func (e ExampleTask) Execute() {

}

func TestNewTask(t *testing.T) {
	pool := NewWorkerBus(1000, 3)
	for i := 0; i < 1000000; i++ {
		pool.SendTask(&ExampleTask{fmt.Sprintf("additional_%d", i+1)})
	}
	pool.Wait()
}

// go test ./tutorial/gorutine/workerpool -run=^# -bench=WorkerPool -benchtime=1000000x -count 5 -cover -v
// Benchmark_WorkerTask-8   	 4746056	       242.8 ns/op	      47 B/op	       2 allocs/op
func Benchmark_WorkerTask(b *testing.B) {
	pool := NewWorkerBus(1000, 3)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool.SendTask(&ExampleTask{fmt.Sprintf("additional_%d", i+1)})
	}
	pool.Wait()
	pool.SendTask(&ExampleTask{"TEST"})
}
