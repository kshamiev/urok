package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, c := context.WithCancel(context.Background())
	done := make(chan struct{})

	wg.Add(1)
	go work1(done)
	wg.Add(1)
	go work2(ctx)

	time.Sleep(time.Second * 5)
	c()
	done <- struct{}{}
	wg.Wait()
}

func work1(done <-chan struct{}) {
	defer func() {
		fmt.Println("work 1 FINISH")
		wg.Done()
	}()
	for {
		// Какая-то работа
		fmt.Println("work 1")
		// Контроль
		select {
		case <-time.After(time.Second):
		case <-done:
			return
		}
	}
}

func work2(ctx context.Context) {
	defer func() {
		fmt.Println("work 2 FINISH")
		wg.Done()
	}()
	for {
		// Какая-то работа
		fmt.Println("work 2")
		// Контроль
		select {
		case <-time.After(time.Second):
		case <-ctx.Done():
			return
		}
	}
}
