package dtbus

import (
	"fmt"
	"log"
	"testing"
)

const (
	countObject = 1000000
)

type ExampleData struct {
	Name string
}

// GOGC=off go test ./tutorial/gorutine/dtbus/. -run=^# -bench=Benchmark_Subscribe -benchtime=1000000x -count 10 -cpu 8
func Benchmark_Subscribe(b *testing.B) {
	b.ReportAllocs()
	Init(100000)

	// подписчики
	for i := 0; i < 1; i++ {
		ch := Gist().Subscribe(&ExampleData{})
		go consumerB(ch)
	}

	b.ResetTimer()
	// отправитель
	for j := 0; j < b.N; j++ {
		Gist().SendData(&ExampleData{Name: fmt.Sprintf("additional_%d", j)})
	}

	Gist().Wait()
}

func consumerB(ch chan interface{}) {
	i := 0
	defer func() {
		if rvr := recover(); rvr != nil {
			log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = Gist().Subscribe(&ExampleData{})
			go consumerB(ch)
		}
	}()
	for obj := range ch {
		_, ok := obj.(*ExampleData)
		if !ok {
			break
		}
		// It`s Work
		// ...
		ch <- true
		i++
	}
	log.Println("full count: ", i)
	close(ch)
}

func Test_Subscribe(t *testing.T) {
	pool := NewDtBus(100000)

	// подписчики
	for i := 0; i < 1; i++ {
		ch := pool.Subscribe(&ExampleData{})
		go consumerT(pool, ch)
	}

	// отправитель
	for i := 0; i < countObject; i++ {
		pool.SendData(&ExampleData{Name: fmt.Sprintf("additional_%d", i+1)})
	}

	pool.Wait()
}

func consumerT(pool *DtBus, ch chan interface{}) {
	i := 0
	defer func() {
		if rvr := recover(); rvr != nil {
			log.Println(fmt.Errorf("%+v", rvr))
			close(ch)
			ch = pool.Subscribe(&ExampleData{})
			go consumerT(pool, ch)
		}
	}()
	for obj := range ch {
		_, ok := obj.(*ExampleData)
		if !ok {
			break
		}
		// It`s Work
		if i == 100000 {
			panic("PANICA")
		}
		// ...
		ch <- true
		i++
	}
	log.Println("full count: ", i)
	close(ch)
}
