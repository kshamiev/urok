package workerdatabus

import (
	"fmt"
	"testing"
	"time"

	"github.com/kshamiev/urok/sample/excel/typs"
)

// GOGC=off go test ./tutorial/gorutine/workerdatabus/. -run=^# -bench=Benchmark_Subscribe -benchtime=100000x -count 3 -cpu 8
func Benchmark_Subscribe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			Subscribe(b)
		}
	}
}

func Subscribe(b *testing.B) {
	pool := NewPool(1000, 3)

	go func() {
		for i := 0; i < 1000000000000; i++ {
			n := fmt.Sprintf("additional_%d", i+1)
			pool.DataSend(&typs.Cargo{Name: n, Amount: 1})
			// time.Sleep(time.Second)
		}
	}()

	// подписчик
	ch1 := pool.DataSubscribe(&typs.Cargo{})
	go func(ch chan interface{}) {
		i := 0
		for obj := range ch {
			if i > 3000 {
				ch <- false
				break
			}
			i++
			o := obj.(*typs.Cargo)
			time.Sleep(time.Millisecond * time.Duration(o.Amount))
			// fmt.Println(o.Name)
			ch <- true
		}
	}(ch1)
	ch2 := pool.DataSubscribe(&typs.Cargo{})
	go func(ch chan interface{}) {
		i := 0
		for obj := range ch {
			if i > 3000 {
				ch <- false
				break
			}
			i++
			o := obj.(*typs.Cargo)
			time.Sleep(time.Millisecond * time.Duration(o.Amount))
			// fmt.Println(o.Name)
			ch <- true
		}
	}(ch2)
	ch3 := pool.DataSubscribe(&typs.Cargo{})
	go func(ch chan interface{}) {
		i := 0
		for obj := range ch {
			if i > 3000 {
				ch <- false
				break
			}
			i++
			o := obj.(*typs.Cargo)
			time.Sleep(time.Millisecond * time.Duration(o.Amount))
			ch <- true
		}
	}(ch3)

	time.Sleep(time.Second * 30)

	pool.Wait()
}
