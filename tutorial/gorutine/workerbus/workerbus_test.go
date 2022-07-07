package workerbus

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/kshamiev/urok/sample/excel/typs"
)

const count = 1000000

// GOGC=off go test ./tutorial/gorutine/workerbus/. -run=^# -bench=Benchmark_Subscribe -benchtime=1000x -count 5 -cpu 8
func Benchmark_Subscribe(b *testing.B) {
	pool := NewWorkerBus(1000, 3)
	for i := 0; i < b.N; i++ {

		// отправитель
		go func() {
			for i := 0; i < count; i++ {
				pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
			}
		}()

		// подписчик
		wg := &sync.WaitGroup{}
		// wg.Add(1)
		pool.SubscribeType(&typs.Cargo{})
		ch := pool.SubscribeType(&typs.Cargo{})
		go consumer(ch, count, wg)
		wg.Wait()

	}
	time.Sleep(time.Minute)
	pool.Wait()
}

func consumer(ch chan interface{}, limitData int, wg *sync.WaitGroup) {
	i := 0
	for obj := range ch {
		_ = obj.(*typs.Cargo)
		// time.Sleep(time.Millisecond * time.Duration(o.Amount))
		i++
		if i == limitData {
			ch <- false
			break
		}
		ch <- true
	}
	fmt.Println(i)
	// wg.Done()
}

func Test_Subscribe(t *testing.T) {
	pool := NewWorkerBus(1000, 3)

	// отправитель
	go func() {
		for i := 0; i < count; i++ {
			pool.SendData(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
		}
	}()

	// подписчик
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := pool.SubscribeType(&typs.Cargo{})
	go consumer(ch, count, wg)
	wg.Wait()

	pool.Wait()
}
