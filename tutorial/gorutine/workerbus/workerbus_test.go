package workerbus

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/kshamiev/urok/sample/excel/typs"
)

const (
	countObject            = 1000000000
	maxLimitConsumerObject = 1000000
)

// GOGC=off go test ./tutorial/gorutine/workerbus/. -run=^# -bench=Benchmark_Subscribe -benchtime=100x -count 5 -cpu 8
func Benchmark_Subscribe(b *testing.B) {
	b.ReportAllocs()
	pool := NewWorkerBus(100000, 3)

	for j := 0; j < b.N; j++ {

		// подписчики
		for i := 0; i < 10; i++ {
			sub := pool.Subscribe(&typs.Cargo{})
			go consumer(sub, int(GenInt(maxLimitConsumerObject)), strconv.Itoa(j)+"-"+strconv.Itoa(i))

		}

		// отправитель
		go func() {
			for i := 0; i < countObject; i++ {
				pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
			}
		}()

	}

	time.Sleep(time.Second)
	pool.Wait()
}

func Test_Subscribe(t *testing.T) {
	pool := NewWorkerBus(100000, 3)

	// подписчики
	for i := 0; i < 10; i++ {
		sub := pool.Subscribe(&typs.Cargo{})
		go consumer(sub, int(GenInt(maxLimitConsumerObject)), strconv.Itoa(i))

	}

	// отправитель
	go func() {
		for i := 0; i < countObject; i++ {
			pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
		}
	}()

	time.Sleep(time.Second)
	pool.Wait()
}

func consumer(sub *Subscribe, limitData int, name string) {
	i := 0
	for obj := range sub.Ch {
		_, ok := obj.(*typs.Cargo)
		if !ok || i == limitData {
			close(sub.Ch)
			break
		}
		// It`s Work
		// ...
		sub.Ch <- true
		i++
	}
	fmt.Println(name+" : consumer work count object: ", i)
}
