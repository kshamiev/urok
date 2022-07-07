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
		for i := 0; i < 10000; i++ {
			n := fmt.Sprintf("additional_%d", i+1)
			pool.DataSend(&typs.Cargo{Name: n, Amount: 1})
			// time.Sleep(time.Second)
		}
	}()

	// подписчик
	ch1 := pool.DataSubscribe(&typs.Cargo{})
	go func(ch chan interface{}) {
		for obj := range ch {
			o := obj.(*typs.Cargo)
			time.Sleep(time.Second * time.Duration(o.Amount))
			// fmt.Println(o.Name)
			ch <- struct{}{}
		}
	}(ch1)
	ch2 := pool.DataSubscribe(&typs.Cargo{})
	go func(ch chan interface{}) {
		for obj := range ch {
			o := obj.(*typs.Cargo)
			time.Sleep(time.Second * time.Duration(o.Amount))
			// fmt.Println(o.Name)
			ch <- struct{}{}
		}
	}(ch2)
	ch3 := pool.DataSubscribe(&typs.Cargo{})
	go func(ch chan interface{}) {
		for obj := range ch {
			o := obj.(*typs.Cargo)
			time.Sleep(time.Second * time.Duration(o.Amount))
			// fmt.Println(o.Name)
			ch <- struct{}{}
		}
	}(ch3)

	time.Sleep(time.Second * 2)
	close(ch1)
	close(ch2)
	close(ch3)

	pool.Wait()
}
