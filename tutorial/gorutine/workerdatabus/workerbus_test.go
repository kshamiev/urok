package workerdatabus

import (
	"fmt"
	"testing"
	"time"

	"github.com/kshamiev/urok/sample/excel/typs"
)

func TestSubscribe(t *testing.T) {
	pool := NewPool(1000, 3)

	go func() {
		for i := 0; i < 1000; i++ {
			n := fmt.Sprintf("additional_%d", i+1)
			pool.DataSend(&typs.Cargo{Name: n})
			// time.Sleep(time.Second)
		}
	}()

	// time.Sleep(time.Second * 3)

	// подписчик
	ch1 := pool.DataSubscribe(&typs.Cargo{})
	go func(ch chan interface{}) {
		for obj := range ch {
			o := obj.(*typs.Cargo)
			fmt.Println(o.Name)
		}
	}(ch1)
	ch2 := pool.DataSubscribe(&typs.Cargo{})
	go func(ch chan interface{}) {
		for obj := range ch {
			o := obj.(*typs.Cargo)
			fmt.Println(o.Name)
		}
	}(ch2)
	ch3 := pool.DataSubscribe(&typs.Cargo{})
	go func(ch chan interface{}) {
		for obj := range ch {
			o := obj.(*typs.Cargo)
			fmt.Println(o.Name)
		}
	}(ch3)

	time.Sleep(time.Millisecond)
	pool.DataUnsubscribeData(&typs.Cargo{}, ch1)
	pool.DataUnsubscribeData(&typs.Cargo{}, ch2)
	pool.DataUnsubscribeData(&typs.Cargo{}, ch3)
	pool.Wait()
}
