package workerbus

import (
	"fmt"
	"sync"
	"testing"

	"github.com/kshamiev/urok/sample/excel/typs"
)

// GOGC=off go test ./tutorial/gorutine/workerdatabus/. -run=^# -bench=Benchmark_Subscribe -benchtime=100000x -count 3 -cpu 8
func Benchmark_Subscribe(b *testing.B) {
	pool := NewPool(1000, 3)
	for i := 0; i < b.N; i++ {

		// отправитель
		go func() {
			for i := 0; i < 1000000; i++ {
				pool.DataSend(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
				// pool.DataSend(&typs.Cargo{Name: "-", Amount: 1})
			}
		}()

		// подписчики
		wg := &sync.WaitGroup{}
		wg.Add(1)
		ch := pool.DataSubscribe(&typs.Cargo{})
		go consumer(ch, 1000000, wg)
		wg.Wait()

	}
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
	wg.Done()
}

func Test_Subscribe(t *testing.T) {
	pool := NewPool(1000, 3)

	// отправитель
	go func() {
		for i := 0; i < 1000000; i++ {
			pool.DataSend(&typs.Cargo{Name: fmt.Sprintf("additional_%d", i+1), Amount: 1})
			// pool.DataSend(&typs.Cargo{Name: "-", Amount: 1})
		}
	}()

	// подписчики
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := pool.DataSubscribe(&typs.Cargo{})
	go consumer(ch, 1000000, wg)
	wg.Wait()

	pool.Wait()
}
