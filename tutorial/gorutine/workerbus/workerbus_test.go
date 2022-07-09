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
		for i := 0; i < 1; i++ {
			sub := pool.Subscribe(&typs.Cargo{})
			go consumer(sub, maxLimitConsumerObject, strconv.Itoa(j)+"-"+strconv.Itoa(i))
		}

		// отправитель
		go func() {
			for i := 0; i < countObject; i++ {
				pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
			}
		}()
	}

	pool.Wait()
}

// GOGC=off go test ./tutorial/gorutine/workerbus/. -run=^# -bench=Benchmark_OneSubscribe -benchtime=1000000x -count 10 -cpu 8
func Benchmark_OneSubscribe(b *testing.B) {
	b.ReportAllocs()
	pool := NewWorkerBus(100000, 3)

	// подписчики
	for i := 0; i < 1; i++ {
		sub := pool.Subscribe(&typs.Cargo{})
		go consumer(sub, maxLimitConsumerObject, strconv.Itoa(i))
	}

	b.ResetTimer()
	// отправитель
	for j := 0; j < b.N; j++ {
		pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", j), Amount: 1})
	}

	pool.Wait()
}

func Test_Subscribe(t *testing.T) {
	pool := NewWorkerBus(100000, 3)

	// подписчики
	for i := 0; i < 1; i++ {
		ch := pool.Subscribe(&typs.Cargo{})
		go consumer(ch, maxLimitConsumerObject, strconv.Itoa(i))
	}

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
