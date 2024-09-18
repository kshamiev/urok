package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	ctx, c := context.WithCancel(context.Background())

	ch1 := make(chan interface{}, 5)
	for i := 0; i < 5; i++ {
		ch1 <- i
	}
	ch2 := make(chan interface{}, 5)
	for i := 0; i < 5; i++ {
		ch2 <- strconv.Itoa(i)
	}
	ch3 := make(chan interface{}, 5)
	for i := 0; i < 5; i++ {
		ch3 <- float64(i)
	}
	chMulti := multiplexer(ctx, ch1, ch2, ch3)

	go func() {
		time.Sleep(time.Second * 5)
		c()
	}()

	for val := range chMulti {
		fmt.Println(val)
	}
}

type TTT interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 | ~string
}

// func multiplexer[T TTT](ctx context.Context, fetchers ...<-chan T) <-chan interface{} {
func multiplexer(ctx context.Context, fetchers ...<-chan interface{}) <-chan interface{} {
	combinedFetcher := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(fetchers))
	for _, f := range fetchers {
		go func(f <-chan interface{}) {
			for {
				select {
				case res := <-f:
					combinedFetcher <- res
				case <-ctx.Done():
					wg.Done()
					return
				}
			}
		}(f)
	}
	go func() {
		wg.Wait()
		close(combinedFetcher)
	}()
	return combinedFetcher
}
