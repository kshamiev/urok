package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	// Генерирует канал, отправляющий целые числа от 0 до 9
	range10 := rangeChannel(done, 10)

	for num := range takeFirstN(done, range10, 5) {
		fmt.Println(num)
	}
}
func takeFirstN(ctx context.Context, dataSource <-chan interface{}, n int) <-chan interface{} {
	// 1
	takeChannel := make(chan interface{})

	// 2
	go func() {
		defer close(takeChannel)

		// 3
		for i := 0; i < n; i++ {
			select {
			case val, ok := <-dataSource:
				if !ok {
					return
				}
				takeChannel <- val
			case <-ctx.Done():
				return
			}
		}
	}()
	return takeChannel
}
