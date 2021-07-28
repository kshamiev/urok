// примеры работы с горутинами
package main

import (
	"fmt"
	// 	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	SampleRuntimeOnce()
}

// пример использование Once & atomic
// выполнение содержимого только один раз (даже если в разных горутинах)
func SampleRuntimeOnce() {
	var onc sync.Once
	var T = int64(345)

	for i := 0; i < 10; i++ {
		go func() {
			atomic.AddInt64(&T, 10)
			fmt.Println(atomic.LoadInt64(&T))
			onc.Do(func() {
				fmt.Println("sync.Once")
			})
		}()
	}
	fmt.Scanln()
	fmt.Print("res: ", T)
}
