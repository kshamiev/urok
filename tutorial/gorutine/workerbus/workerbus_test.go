package workerbus

import (
	"fmt"
	"testing"
	"time"

	"github.com/kshamiev/urok/sample/excel/typs"
)

const count = 1000000

// GOGC=off go test ./tutorial/gorutine/workerbus/. -run=^# -bench=Benchmark_Subscribe -benchtime=1000x -count 5 -cpu 8
func Benchmark_Subscribe(b *testing.B) {
	b.ReportAllocs()
	pool := NewWorkerBus(1000, 3)

	// подписчик
	sub := pool.Subscribe(&typs.Cargo{})
	go consumer(sub, count-10)

	// отправитель
	// go func() {
	for i := 0; i < count; i++ {
		pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
	}
	// }()

	// time.Sleep(time.Second)

	pool.Wait()
}

func Test_Subscribe(t *testing.T) {
	pool := NewWorkerBus(1000000, 3)

	// подписчики
	for i := 0; i < 10; i++ {
		sub := pool.Subscribe(&typs.Cargo{})
		go consumer(sub, int(GenInt(10000000)))

	}

	// отправитель
	go func() {
		for i := 0; i < count; i++ {
			pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
		}
	}()

	time.Sleep(time.Second * 2)
	pool.Wait()
}

func consumer(sub *Subscribe, limitData int) {
	i := 0
	for obj := range sub.Ch {
		_ = obj.(*typs.Cargo)
		i++
		if i == limitData {
			close(sub.Ch)
			fmt.Println()
			fmt.Println("consumer finish (limit or condition)")
			break
		}
		sub.Ch <- true
	}
	fmt.Println("consumer work count object: ", i)
}
