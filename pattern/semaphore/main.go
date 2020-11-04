// Семафор это решение которое контролирует количество процессов - горутин
// работающих с каким-либо ресурсом или выполняющих какую-то иную работу параллельно.
package main

import (
	"fmt"
	"math/rand"
	"time"

	"urok/pattern/semaphore/sem"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var n = 100
	var m = 1000
	var data = make(map[string][]int64)
	for i := 0; i < n; i++ {
		row := make([]int64, 0, m)
		for j := 0; j < m; j++ {
			row = append(row, rand.Int63())
		}
		data[fmt.Sprintf("%d", rand.Int63())] = row
	}
	result := semaphore(data, rand.Intn(9)+1)
	for k, v := range result {
		fmt.Printf("k:%s v:%d\n", k, v)
	}
}

func semaphore(src map[string][]int64, threads int) map[string]int64 {
	s := sem.NewPool(threads)
	return s.Work(src)
}
