package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx)
	ctx = context.TODO()
	fmt.Println(ctx)

	ctx = context.WithValue(ctx, "key", "test")
	val := ctx.Value("key")
	ch := ctx.Done()
	err := ctx.Err()
	timeDedline, ok := ctx.Deadline()
	fmt.Println(val, ch, err, timeDedline, ok)
	fmt.Println()
	fmt.Println(runtime.GoroutineProfile(nil))

	testWithCancel(ctx)
	testWithDeadline(ctx)
	testWithTimeout(ctx)

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second)
		fmt.Println(runtime.GoroutineProfile(nil))
	}
}

func testWithCancel(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("WithCancel cancel")
		cancel()
	}()
	done := <-ctx.Done()
	fmt.Printf("func WithCancel done: %v\n\n", done)
}

func testWithDeadline(ctx context.Context) {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("WithDeadline cancel")
		cancel()
	}()
	done := <-ctx.Done()
	fmt.Printf("func WithDeadline done: %v\n\n", done)
}

func testWithTimeout(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	go func() {
		time.Sleep(time.Second * 15)
		fmt.Println("WithTimeout cancel")
		cancel()
	}()
	done := <-ctx.Done()
	fmt.Printf("func WithTimeout done: %v\n\n", done)
}
