package workerbus

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/kshamiev/urok/sample/excel/typs"
)

const (
	countObject            = 1000000
	maxLimitConsumerObject = 1000000
)

// GOGC=off go test ./tutorial/gorutine/workerbus/. -run=^# -bench=Benchmark_Subscribe -benchtime=100x -count 1 -cpu 8
func Benchmark_Subscribe(b *testing.B) {
	b.ReportAllocs()
	pool := NewWorkerBus(100000, 3)
	b.ResetTimer()

	for j := 0; j < b.N; j++ {
		// подписчики
		sub := pool.Subscribe(&typs.Cargo{})
		go consumer(sub, int(GenInt(maxLimitConsumerObject)), strconv.Itoa(j))

		// отправитель
		for i := 0; i < countObject; i++ {
			pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
		}
	}

	pool.Wait()
}

func Test_Subscribe(t *testing.T) {
	pool := NewWorkerBus(100000, 3)

	// подписчики
	i := 0
	ch := pool.Subscribe(&typs.Cargo{})
	go consumer(ch, int(GenInt(maxLimitConsumerObject)), strconv.Itoa(i))

	// отправитель
	for i := 0; i < countObject; i++ {
		pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
	}

	pool.Wait()
}

func consumer(ch chan interface{}, limitData int, name string) {
	i := 0
	for obj := range ch {
		_, ok := obj.(*typs.Cargo)
		if !ok || i == limitData {
			close(ch)
			break
		}
		// It`s Work
		// ...
		ch <- true
		i++
	}
}
