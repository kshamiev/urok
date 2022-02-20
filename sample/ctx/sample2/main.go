package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	ctx, _ := context.WithCancel(context.Background())

	go test1(ctx)
	go test1(ctx)
	go test1(ctx)

	ctx1, _ := context.WithTimeout(ctx, 10*time.Second)
	go func() {
		<-ctx1.Done()
		fmt.Println("main done")
	}()

	var tt int
	fmt.Scan(&tt)
	fmt.Println(tt)
}

func test1(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("test1 done")
}
